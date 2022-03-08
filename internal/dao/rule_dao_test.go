package dao

import (
	"alert/internal/evaluator/rules"
	"alert/internal/model"
	"testing"
)

var ruleDao RuleDao

func init() {
	ruleDao = *NewRuleDao()
}

func TestAddNormalRule(t *testing.T) {
	rule := model.Rule{}
	rule.Code = "for 13 room half of turnover recent 3 min can/'t under 20"
	rule.Name = "3分钟营业额的一半不能低于20"
	rule.RoomId = 13
	rule.Type = model.NORMALRULE
	rule.Expr = "index[half of turnover recent 3 min] >= 20"
	ruleNode, _ := rules.ToNormalRuleExpr(rule.Expr)
	rule.Serialized = ruleNode.ToJson()
	_, err := ruleDao.AddRule(&rule)
	if err != nil {
		t.Error(err.Error())
	}

	rule = model.Rule{}
	rule.Code = "turnover equal 200"
	rule.Name = "营业额等于200"
	rule.RoomId = 13
	rule.Type = model.NORMALRULE
	rule.Expr = "index[turnover] = 200"
	ruleNode, _ = rules.ToNormalRuleExpr(rule.Expr)
	rule.Serialized = ruleNode.ToJson()
	_, err = ruleDao.AddRule(&rule)
	if err != nil {
		t.Error(err.Error())
	}
}
func TestAddCompleteRule(t *testing.T) {
	rule := model.Rule{}
	rule.Code = "for 13 room half of turnover recent 3 min can/'t under 20 or equal 200"
	rule.Name = "3分钟营业额的一半不能低于20或等于营业额200"
	rule.RoomId = 18
	rule.Type = model.COMPLEXRULE
	rule.Expr = "rule[for 13 room half of turnover recent 3 min can/'t under 20]|rule[turnover equal 200]"
	completeRule, _ := rules.ToCompleteRuleExpr(rule.Expr)
	rule.Serialized = completeRule.ToJson()
	_, err := ruleDao.AddRule(&rule)
	if err != nil {
		t.Error(err.Error())
	}
}
func TestDeleteRule(t *testing.T) {
	_, err := ruleDao.DeleteRuleByID(3)
	if err != nil {
		t.Error(err.Error())
	}
}
func TestDeleteRuleByCode(t *testing.T) {
	_, err := ruleDao.DeleteRuleByCode("for 13 room half of turnover recent 3 min can/'t under 20 or equal 200")
	if err != nil {
		t.Error(err.Error())
	}
}
func TestSelectRule(t *testing.T) {
	rule, err := ruleDao.SelectRuleByID(3)
	if err != nil {
		t.Error(err.Error())
	}
	print(rule.Code)
}
