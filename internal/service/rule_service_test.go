package service

import (
	"testing"
)

var ruleServcie RuleService

func init() {
	ruleServcie = *NewRuleService()
}

func TestAddRule(t *testing.T) {
	ruleServcie.AddRule(0, "营业额总和大于80", "sum of turnover up 80", true, "index[turnover] > 80")                //T
	ruleServcie.AddRule(0, "营业额总和小于120", "sum of turnover under 120", true, "index[turnover] < 120")          //T
	ruleServcie.AddRule(0, "营业额总和大于等于100", "sum of turnover up or equal 100", true, "index[turnover] >= 100") //T
	ruleServcie.AddRule(0, "营业额总和大于100", "sum of turnover up  100", true, "index[turnover] > 100")            //F
	ruleServcie.AddRule(0, "营业额总和大于120", "sum of turnover up 120", true, "index[turnover] > 120")             //F

	ruleStr := "rule[sum of turnover up 80] | rule[sum of turnover under 120]"
	ruleServcie.AddRule(0, "营业额总和大于80或小于120", "sum of turnover up 80 or under 120", false, ruleStr) //T
	ruleStr = "rule[sum of turnover up 80] ^ rule[sum of turnover under 120]"
	ruleServcie.AddRule(0, "测试0", "test0", false, ruleStr) //F
	ruleStr = "((rule[sum of turnover up 80]&rule[sum of turnover under 120])|(rule[sum of turnover under 120]^rule[sum of turnover up  100]))&(rule[sum of turnover up  100])"
	ruleServcie.AddRule(0, "测试1", "test1", false, ruleStr) //F
	ruleStr = "(rule[sum of turnover up 80]&rule[sum of turnover under 120])|(rule[sum of turnover under 120]^rule[sum of turnover up  100])&(rule[sum of turnover up  100])"
	ruleServcie.AddRule(0, "测试2", "test2", false, ruleStr) //T
	ruleStr = "rule[test0]^rule[test2]"
	ruleServcie.AddRule(0, "测试3", "test3", false, ruleStr) //T
}
func TestAddRuleFault(t *testing.T) {
	ruleStr := "rule[test4]|rule[test3]"
	_, err := ruleServcie.AddRule(0, "测试4", "test4", false, ruleStr)
	if err == nil {
		t.Error("未捕获错误")
	}
}
func TestCheckRule(t *testing.T) {
	var i uint
	want := []bool{true, true, true, false, false, true, false, false, true, true}
	for i = 11; i <= 20; i++ {
		rule, _ := ruleServcie.SelectRuleById(i)
		r, err := ruleServcie.CheckRule(rule.Code)
		if err != nil {
			print(err.Error())
		}
		if r != want[i-11] {
			t.Error()
		}
	}
}
