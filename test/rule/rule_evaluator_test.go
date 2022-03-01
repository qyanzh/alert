package rule

import (
	"alert/internal/evaluator"
	"strconv"
	"testing"
)

func TestToNormalRuleExpr(t *testing.T) {
	expr := "index[turning] <= 200"
	node := evaluator.ToNormalRuleExpr(expr)
	println("code:" + node.IndexCode)
	println("op:" + node.Op)
	println("num:" + strconv.FormatFloat(node.Number, 'e', 16, 64))
}
