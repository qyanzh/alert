package rule

import (
	"alert/internal/evaluator"
	"strconv"
	"testing"
)

func TestToNormalRuleExpr(t *testing.T) {
	expr := "index[turnover] <= 200"
	node := evaluator.ToNormalRuleExpr(expr)
	println("id:" + strconv.FormatUint(uint64(node.IndexId), 10))
	println("op:" + node.Op)
	println("num:" + strconv.FormatFloat(node.Number, 'e', 16, 64))
}
