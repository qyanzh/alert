package service

import "alert/internal/dao"

type RuleService struct {
	alertDao dao.RuleDao
}

func NewRuleService() *RuleService {
	return &RuleService{alertDao: *dao.NewRuleDao()}
}
