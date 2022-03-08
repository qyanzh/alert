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

func (service *RuleService) SelectRule(code string) (*model.Rule, error) {
	return service.ruleDao.SelectRuleByCode(code)
}

func (service *RuleService) SelectAllRules() (*[]model.Rule, error) {
	return service.ruleDao.SelectAllRules()
}

func (service *RuleService) SelectRuleById(id uint) (*model.Rule, error) {
	return service.ruleDao.SelectRuleByID(id)
}

//ruleType 为true是normalrule
func (service *RuleService) AddRule(roomId uint, name string, code string, ruleType bool, content string) (*model.Rule, error) {
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
	_, err := service.ruleDao.AddRule(&rule)
	if err != nil {
		return &rule, err
	}
	return service.ruleDao.SelectRuleByCode(code)
}

func (service *RuleService) DeleteRule(code string) error {
	_, err := service.ruleDao.DeleteRuleByCode(code)
	return err
}

func (service *RuleService) UpdateRule(rule model.Rule) error {
	_, err := service.ruleDao.UpdateRule(&rule)
	return err
}

func (service *RuleService) checkANum(index float64, op string, num float64) (bool, error) {
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
		return num > index, nil
	case "!=":
		return math.Abs(index-num) >= 1e-8, nil
	default:
		return false, errors.New("未知符号")
	}
}

func (service *RuleService) checkNormalRule(normalRule *rules.NormalRule, roomId uint) (bool, error) {
	//indexnum:=
	indexNum := 0.0
	return service.checkANum(indexNum, normalRule.Op, normalRule.Number)
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
func (service *RuleService) getRuleById(ruleId uint) (*model.Rule, error) {
	return service.ruleDao.SelectRuleByID(ruleId)
}

func (service *RuleService) getAllNum(completeRule *rules.CompleteRule, ruleMap map[uint]*model.Rule, ids *[]uint) error {
	var err error
	for _, value := range *completeRule {
		if value.Type == rules.RULENODE {
			ruleId, _ := value.Content.(uint)
			rule, ok := ruleMap[ruleId]
			if !ok {
				rule, err = service.getRuleById(ruleId)
				if err != nil {
					return err
				}
				ruleMap[ruleId] = rule
			}
			if rule.Type == model.NORMALRULE {
				*ids = append(*ids, rules.GetNormalRule(rule.Serialized).IndexId)
			} else {
				err = service.getAllNum(rules.GetCompleteRule(rule.Serialized), ruleMap, ids)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (service *RuleService) checkAllNum(completeRule *rules.CompleteRule, indexMap map[uint]float64, ruleMap map[uint]*model.Rule) (bool, error) {
	s := make(boolStack, 0)
	var err error
	for _, value := range *completeRule {
		if value.Type == rules.RULENODE {
			rule := ruleMap[value.Content.(uint)]
			var r bool
			if rule.Type == model.NORMALRULE {
				normalRule := rules.GetNormalRule(rule.Serialized)
				r, err = service.checkANum(indexMap[rule.Id], normalRule.Op, normalRule.Number)
				if err != nil {
					return false, err
				}
			} else {
				completeRule := rules.GetCompleteRule(rule.Serialized)
				r, err = service.checkAllNum(completeRule, indexMap, ruleMap)
				if err != nil {
					return false, err
				}
			}
			s.Push(r)
		} else if value.Type == rules.RULEOP {
			if len(s) < 2 {
				return false, errors.New("请检查语法")
			}
			op := value.Content.(rune)
			r1 := s.Top()
			s.Pop()
			r2 := s.Top()
			s.Pop()
			if op == '&' {
				s.Push(r1 && r2)
			} else if op == '|' {
				s.Push(r1 || r2)
			} else if op == '^' {
				s.Push((r1 || r2) && !(r1 && r2))
			} else {
				return false, errors.New("请检查语法")
			}
		}
	}
	if len(s) != 1 {
		return false, errors.New("请检查语法")
	}
	return s.Top(), nil
}

func (service *RuleService) checkCompleteRule(completeRule *rules.CompleteRule, roomId uint) (bool, error) {

	ids := make([]uint, 0)
	ruleMap := make(map[uint]*model.Rule, 0)
	err := service.getAllNum(completeRule, ruleMap, &ids)
	if err != nil {
		return false, err
	}
	indexMap := make(map[uint]float64, 0)
	return service.checkAllNum(completeRule, indexMap, ruleMap)
}

func (service *RuleService) CheckRule(code string) (bool, error) {
	rule, err := service.SelectRule(code)
	if err != nil {
		return false, err
	}
	if rule.Type == model.NORMALRULE {
		return service.checkNormalRule(rules.GetNormalRule(rule.Serialized), rule.RoomId)
	} else {
		return service.checkCompleteRule(rules.GetCompleteRule(rule.Serialized), rule.RoomId)
	}
}

func (service *RuleService) CheckRuleWithId(id uint) (bool, error) {
	rule, err := service.SelectRuleById(id)
	if err != nil {
		return false, nil
	}
	if rule.Type == model.NORMALRULE {
		return service.checkNormalRule(rules.GetNormalRule(rule.Serialized), rule.RoomId)
	} else {
		return service.checkCompleteRule(rules.GetCompleteRule(rule.Serialized), rule.RoomId)
	}
}
