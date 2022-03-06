package service

import (
	"alert/internal/dao"
	"alert/internal/evaluator"
	"alert/internal/model"
	"errors"
)

type RuleService struct {
	ruleDao dao.RuleDao
}

func NewRuleService() *RuleService {
	return &RuleService{ruleDao: *dao.NewRuleDao()}
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
	rule := model.Rule{Code: code, Name: name, RoomId: roomId, Expr: content}
	if ruleType {
		rule.Type = model.Normal_Rule
		ruleNode, _ := evaluator.ToNormalRuleExpr(rule.Expr)
		rule.Serialized = ruleNode.ToJson()
	} else {
		rule.Type = model.Complex_Rule
		ruleNode, _ := evaluator.ToCompleteRuleExpr(rule.Expr)
		rule.Serialized = ruleNode.ToJson()
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

func (service *RuleService) checkNormalRule(normalRule *evaluator.NormalRule, roomId uint) (bool, error) {
	return false, nil
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

func (service *RuleService) checkCompleteRule(completeRule *evaluator.CompleteRule, roomId uint) (bool, error) {
	s := make(boolStack, 0)
	for _, value := range *completeRule {
		if value.Type == evaluator.RuleNode {
			rule, err := service.getRuleById(value.Content.(uint))
			if err != nil {
				return false, err
			}
			r, err := service.checkNormalRule(evaluator.GetNormalRule(rule.Serialized), roomId)
			if err != nil {
				return false, err
			}
			s.Push(r)
		} else if value.Type == evaluator.RuleOp {
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

func (service *RuleService) CheckRule(code string) (bool, error) {
	rule, err := service.SelectRule(code)
	if err != nil {
		return false, err
	}
	if rule.Type == model.Normal_Rule {
		return service.checkNormalRule(evaluator.GetNormalRule(rule.Serialized), rule.RoomId)
	} else {
		return service.checkCompleteRule(evaluator.GetCompleteRule(rule.Serialized), rule.RoomId)
	}
}
