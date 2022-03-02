package service

import (
	"alert/internal/dao"
	"alert/internal/model"
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
