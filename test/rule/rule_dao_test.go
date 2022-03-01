package rule

import (
	"alert/internal/dao"
	"alert/internal/evaluator"
	"alert/internal/model"
	"testing"
)

var ruleDao dao.RuleDao

func init() {
	ruleDao = *dao.NewRuleDao()
}

func TestAddNormalRule(t *testing.T) {
	rule := model.Rule{}
	rule.Code = "for 13 room half of turnover recent 3 min can/'t under 20"
	rule.Name = "3分钟营业额的一半不能低于20"
	rule.RoomId = 13
	rule.Type = model.Normal_Rule
	rule.Expr = "index[half of turnover recent 3 min] >= 20"
	ruleNode, _ := evaluator.ToNormalRuleExpr(rule.Expr)
	rule.Serialized = ruleNode.ToJson()
	ruleDao.AddRule(&rule)

	rule = model.Rule{}
	rule.Code = "turnover equal 200"
	rule.Name = "营业额等于200"
	rule.RoomId = 13
	rule.Type = model.Normal_Rule
	rule.Expr = "index[turnover] = 200"
	ruleNode, _ = evaluator.ToNormalRuleExpr(rule.Expr)
	rule.Serialized = ruleNode.ToJson()
	ruleDao.AddRule(&rule)
}

func TestDeleteRule(t *testing.T) {

}
