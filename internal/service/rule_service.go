package service

import (
	"alert/internal/dao"
	"alert/internal/evaluator/rules"
	"alert/internal/model"
	"errors"
	"math"
)

type RuleService struct {
	ruleDao      dao.RuleDao
	indexService IndexService
}

func NewRuleService() *RuleService {
	return &RuleService{ruleDao: *dao.NewRuleDao(), indexService: *NewIndexService()}
}

func (rs *RuleService) SelectRule(code string) (*model.Rule, error) {
	return rs.ruleDao.SelectRuleByCode(code)
}

func (rs *RuleService) SelectAllRules() (*[]model.Rule, error) {
	return rs.ruleDao.SelectAllRules()
}

func (rs *RuleService) SelectRuleById(id uint) (*model.Rule, error) {
	return rs.ruleDao.SelectRuleByID(id)
}

//ruleType 为true是normalrule
func (rs *RuleService) AddRule(roomId uint, name string, code string, ruleType bool, content string) (*model.Rule, error) {
	rule := model.Rule{
		Code:   code,
		Name:   name,
		RoomId: roomId,
		Expr:   content,
	}
	if ruleType {
		rule.Type = model.NORMALRULE
		ruleNode, err := rules.ToNormalRuleExpr(rule.Expr)
		rule.Serialized = ruleNode.ToJson()
		if err != nil {
			return nil, err
		}
	} else {
		rule.Type = model.COMPLEXRULE
		ruleNode, err := rules.ToCompleteRuleExpr(rule.Expr)
		rule.Serialized = ruleNode.ToJson()
		if err != nil {
			return nil, err
		}
	}
	_, err := rs.ruleDao.AddRule(&rule)
	if err != nil {
		return &rule, err
	}
	return rs.ruleDao.SelectRuleByCode(code)
}

func (rs *RuleService) DeleteRule(code string) error {
	_, err := rs.ruleDao.DeleteRuleByCode(code)
	return err
}
func (rs *RuleService) EvaluatorRule(rule *model.Rule) ([]byte, error) {
	if rule.Type == model.NORMALRULE {
		ruleNode, err := rules.ToNormalRuleExpr(rule.Expr)
		if err != nil {
			return nil, err
		}
		return ruleNode.ToJson(), err
	} else {
		ruleNode, err := rules.ToCompleteRuleExpr(rule.Expr)
		if err != nil {
			return nil, err
		}
		return ruleNode.ToJson(), err
	}
}
func (rs *RuleService) UpdateRule(rule *model.Rule) error {
	_, err := rs.ruleDao.UpdateRule(rule)
	return err
}

func (rs *RuleService) checkANum(index float64, op string, num float64) (bool, error) {
	switch op {
	case "=":
		return math.Abs(index-num) < 1e-8, nil
	case "<=":
		return index-num < 1e-8, nil
	case "<":
		return index < num, nil
	case ">=":
		return num-index < 1e-8, nil
	case ">":
		return index > num, nil
	case "!=":
		return math.Abs(index-num) >= 1e-8, nil
	default:
		return false, errors.New("未知符号")
	}
}

func (rs *RuleService) checkNormalRule(normalRule *rules.NormalRule, roomId uint) (bool, map[uint]float64, error) {
	id := []uint{normalRule.IndexId}
	indexNum, _ := rs.indexService.SelectIndexValuesByIDsAndRoomID(id, roomId)
	//if err!=nil{
	//	return false,err
	//}
	rb, err := rs.checkANum(indexNum[id[0]], normalRule.Op, normalRule.Number)
	return rb, indexNum, err
}

type boolStack []bool

func (s *boolStack) Top() bool {
	return (*s)[len(*s)-1]
}
func (s *boolStack) Push(num bool) {
	(*s) = append(*s, num)
}
func (s *boolStack) Pop() {
	*s = (*s)[:len(*s)-1]
}
func (rs *RuleService) getRuleById(ruleId uint) (*model.Rule, error) {
	return rs.ruleDao.SelectRuleByID(ruleId)
}

const MAXDEPTH = 10

func (rs *RuleService) getAllNum(completeRule *rules.CompleteRule, ruleMap map[uint]*model.Rule, ids *[]uint, nowDepth uint) error {
	if nowDepth == MAXDEPTH {
		return errors.New("可能出现无限递归")
	}
	var err error
	for _, value := range *completeRule {
		if value.Type == rules.RULENODE {
			ruleId := value.Content
			rule, ok := ruleMap[ruleId]
			if !ok {
				rule, err = rs.getRuleById(ruleId)
				if err != nil {
					return err
				}
				ruleMap[ruleId] = rule
			}
			if rule.Type == model.NORMALRULE {
				*ids = append(*ids, rules.GetNormalRule(rule.Serialized).IndexId)
			} else {
				err = rs.getAllNum(rules.GetCompleteRule(rule.Serialized), ruleMap, ids, nowDepth+1)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (rs *RuleService) checkAllNum(completeRule *rules.CompleteRule, indexMap map[uint]float64, ruleMap map[uint]*model.Rule) (bool, error) {
	s := make(boolStack, 0)
	var err error
	for _, value := range *completeRule {
		if value.Type == rules.RULENODE {
			rule := ruleMap[value.Content]
			var r bool
			if rule.Type == model.NORMALRULE {
				normalRule := rules.GetNormalRule(rule.Serialized)
				r, err = rs.checkANum(indexMap[normalRule.IndexId], normalRule.Op, normalRule.Number)
				if err != nil {
					return false, err
				}
			} else {
				completeRule := rules.GetCompleteRule(rule.Serialized)
				r, err = rs.checkAllNum(completeRule, indexMap, ruleMap)
				if err != nil {
					return false, err
				}
			}
			s.Push(r)
		} else if value.Type == rules.RULEOP {
			if len(s) < 2 {
				return false, errors.New("请检查语法")
			}
			op := value.Content
			r1 := s.Top()
			s.Pop()
			r2 := s.Top()
			s.Pop()
			if op == '&' {
				s.Push(r1 && r2)
			} else if op == '|' {
				s.Push(r1 || r2)
			} else if op == '^' || op == '!' {
				s.Push((r1 || r2) && !(r1 && r2))
			} else {
				return false, errors.New("请检查语法")
			}
		} else if value.Type == rules.RULENUMBER {
			if value.Type == 1 {
				s.Push(true)
			} else {
				s.Push(false)
			}
		}
	}
	if len(s) != 1 {
		return false, errors.New("请检查语法")
	}
	return s.Top(), nil
}

func (rs *RuleService) checkCompleteRule(completeRule *rules.CompleteRule, roomId uint) (bool, map[uint]float64, error) {

	ids := make([]uint, 0)
	ruleMap := make(map[uint]*model.Rule, 0)
	err := rs.getAllNum(completeRule, ruleMap, &ids, 1)
	if err != nil {
		return false, nil, err
	}
	indexMap, err := rs.indexService.SelectIndexValuesByIDsAndRoomID(ids, roomId)
	//if err != nil {
	//	return false, err
	//}
	rb, err := rs.checkAllNum(completeRule, indexMap, ruleMap)
	return rb, indexMap, err
}

func (rs *RuleService) CheckRule(code string) (bool, map[uint]float64, error) {
	rule, err := rs.SelectRule(code)
	if err != nil {
		return false, nil, err
	}
	if rule.Type == model.NORMALRULE {
		return rs.checkNormalRule(rules.GetNormalRule(rule.Serialized), rule.RoomId)
	} else {
		return rs.checkCompleteRule(rules.GetCompleteRule(rule.Serialized), rule.RoomId)
	}
}

func (rs *RuleService) CheckRuleWithId(id uint) (bool, map[uint]float64, error) {
	rule, err := rs.SelectRuleById(id)
	if err != nil {
		return false, nil, err
	}
	if rule.Type == model.NORMALRULE {
		return rs.checkNormalRule(rules.GetNormalRule(rule.Serialized), rule.RoomId)
	} else {
		return rs.checkCompleteRule(rules.GetCompleteRule(rule.Serialized), rule.RoomId)
	}
}
