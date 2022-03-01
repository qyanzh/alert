package rule

import (
	"alert/internal/evaluator"
	"strconv"
	"testing"
)

func TestToNormalRuleExpr(t *testing.T) {
	expr := "index[turnover] <= 200"
	node, _ := evaluator.ToNormalRuleExpr(expr)
	println("id:" + strconv.FormatUint(uint64(node.IndexId), 10))
	println("op:" + node.Op)
	println("num:" + strconv.FormatFloat(node.Number, 'e', 16, 64))
}

func TestToComplexRuleExpr(t *testing.T) {
	expr := "rule[for 13 room half of turnover recent 3 min can/'t under 20]&rule[for 13 room half of turnover recent 3 min can/'t under 20]"
	nodes, _ := evaluator.ToCompleteRuleExpr(expr)
	for _, value := range *nodes {
		value.PrintContent()
		print(" ")
	}
}
