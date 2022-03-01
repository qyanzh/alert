package evaluator

import (
	"encoding/json"
	"log"
	"strconv"
)

type NormalRuleNode struct {
	IndexCode string
	Number    float64
	Op        string
}

func (r *NormalRuleNode) ToJson() []byte {
	s, err := json.Marshal(r)
	if err != nil {
		log.Panicln(err)
	}
	return s
}

func isCompareOp(c uint8) bool {
	if c == '>' || c == '<' || c == '=' || c == '!' {
		return true
	}
	return false
}

//从nowst开始找下一个code
func getNextCode(expr string, nowst int) (st int, ed int) {
	for st = nowst; st < len(expr); st++ {
		if expr[st] == '[' {
			st++
			break
		}
	}
	for ed = st; ed < len(expr); ed++ {
		if expr[ed] == ']' {
			break
		}
	}
	return
}

//从nowst开始找下一个运算符
func getNextOp(expr string, nowst int) (st int, ed int) {
	for st = nowst; st < len(expr); st++ {
		if isCompareOp(expr[st]) {
			break
		}
	}
	for ed = st; ed < len(expr); ed++ {
		if !isCompareOp(expr[ed]) {
			break
		}
	}
	return
}

//从nowst开始找下一个数字
func getNextNum(expr string, nowst int) (st int, ed int) {
	for st = nowst; st < len(expr); st++ {
		if expr[st] <= '9' && expr[st] >= '0' {
			break
		}
	}
	for ed = st; ed < len(expr); ed++ {
		if expr[st] > '9' || expr[st] < '0' {
			break
		}
	}
	return
}
func ToNormalRuleExpr(expr string) NormalRuleNode {
	node := NormalRuleNode{}
	st, ed := getNextCode(expr, 0)
	node.IndexCode = expr[st:ed]
	st, ed = getNextOp(expr, ed)
	node.Op = expr[st:ed]
	st, ed = getNextNum(expr, ed)
	node.Number, _ = strconv.ParseFloat(expr[st:ed], 64)
	return node
}
