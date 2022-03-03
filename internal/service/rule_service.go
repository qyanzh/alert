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

func (service *RuleService) AddRule(roomId uint, name string, code string, ruleType bool, content string) (*model.Rule, error) {
	var nowType model.RuleType
	if ruleType {
		nowType = model.Normal_Rule
	} else {
		nowType = model.Complex_Rule
	}
	rule := model.Rule{Code: code, Name: name, RoomId: roomId, Type: nowType, Expr: content}
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
			r, err := service.checkNormalRule(evaluator.GetNormalRule(rule.Serialized), rule.RoomId)
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
