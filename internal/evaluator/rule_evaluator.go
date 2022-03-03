package evaluator

import (
	"alert/internal/dao"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

var indexDao dao.IndexDao
var ruleDao dao.RuleDao

func init() {
	indexDao = *dao.NewIndexDao()
	ruleDao = *dao.NewRuleDao()
}

type NormalRuleNode struct {
	IndexId uint
	Number  float64
	Op      string
}

type CompleteType int8

const (
	RuleOp CompleteType = iota
	RuleNumber
	RuleNode
)

type CompleteNode struct {
	Type    CompleteType
	Content interface{}
}

func (r *CompleteNode) PrintContent() {
	if r.Type == RuleNode {
		x, _ := r.Content.(uint)
		print(x)
	} else if r.Type == RuleOp {
		x, _ := r.Content.(rune)
		fmt.Printf("%c", x)
	} else if r.Type == RuleNumber {
		x, _ := r.Content.(float64)
		print(x)
	}
}
func (r *NormalRuleNode) ToJson() []byte {
	s, err := json.Marshal(r)
	if err != nil {
		log.Panicln(err)
	}
	return s
}
func GetNormalRule(b []byte) *NormalRule {
	rule := NormalRule{}
	err := json.Unmarshal(b, &rule)
	if err != nil {
		log.Panicln(err)
	}
	return &rule
}
func (r *CompleteRule) ToJson() []byte {
	s, err := json.Marshal(r)
	if err != nil {
		log.Panicln(err)
	}
	return s
}
func GetCompleteRule(b []byte) *CompleteRule {
	rule := CompleteRule{}
	err := json.Unmarshal(b, &rule)
	if err != nil {
		log.Panicln(err)
	}
	return &rule
}

//普通节点的处理
type NormalRule NormalRuleNode

var oplist = []uint8{'>', '<', '=', '!', '&', '|', '^', '(', ')'}

func isCompareOp(c uint8) bool {
	for _, value := range oplist {
		if c == value {
			return true
		}
	}
	return false
}

//从nowst开始找下一个code
type strType uint8

const (
	codeType strType = iota
	opType
	numType
	other
)

func getNext(expr string, nowst int, isSingleOp bool) (st int, ed int, nowType strType) {
	nowType = other
	for st = nowst; st < len(expr); st++ {
		if expr[st] == '[' {
			st++
			nowType = codeType
			for ed = st; ed < len(expr); ed++ {
				if expr[ed] == ']' {
					break
				}
			}
			break
		}
		if expr[st] <= '9' && expr[st] >= '0' {
			nowType = numType
			for ed = st; ed < len(expr); ed++ {
				if expr[st] > '9' || expr[st] < '0' {
					break
				}
			}

			break
		}
		if isCompareOp(expr[st]) {
			nowType = opType
			if isSingleOp {
				ed = st + 1
				break
			}
			for ed = st; ed < len(expr); ed++ {
				if !isCompareOp(expr[ed]) {
					break
				}
			}
			break
		}
	}

	return
}

func ToNormalRuleExpr(expr string) (NormalRuleNode, error) {
	node := NormalRuleNode{}
	st, ed, nowType := getNext(expr, 0, false)
	if nowType != codeType {
		return node, errors.New("语法错误")
	}
	node.IndexId = indexDao.SelectIndexByCode(expr[st:ed]).ID
	st, ed, nowType = getNext(expr, ed, false)
	if nowType != opType {
		return node, errors.New("语法错误")
	}
	node.Op = expr[st:ed]
	st, ed, nowType = getNext(expr, ed, false)
	if nowType != numType {
		return node, errors.New("语法错误")
	}
	node.Number, _ = strconv.ParseFloat(expr[st:ed], 64)
	return node, nil
}

//复杂节点的处理

type CompleteRule []CompleteNode

var logicOpPriority = map[rune]int8{
	'&': 3,
	'^': 2,
	'|': 1,
}

func logicOpGE(o1, o2 rune) bool {
	return logicOpPriority[o1]-logicOpPriority[o2] >= 0
}
func isLogicOp(c rune) bool {
	return c == '&' || c == '|' || c == '^'
}
func ToCompleteRuleExpr(expr string) (*CompleteRule, error) {
	nodes := make(CompleteRule, 0)
	opStack := make(stack, 0)
	var st, ed int
	var nowType strType
	for ed < len(expr) {
		st, ed, nowType = getNext(expr, ed, true)
		if st >= len(expr) {
			break
		}
		if nowType == codeType {
			index, err := ruleDao.SelectRuleByCode(expr[st:ed])
			if err != nil {
				return &nodes, err
			}
			nodes = append(nodes, CompleteNode{Type: RuleNode, Content: index.Id})
		} else if nowType == numType {
			return &nodes, errors.New("请检查语法")
		} else if nowType == opType {
			r := rune(expr[st])
			if isLogicOp(r) {
				// 弹出栈中优先级>=当前运算符的运算符
				for top := opStack.peek(); top != 0 && top != '(' && logicOpGE(top, r); top = opStack.peek() {
					opStack, _ = opStack.pop()
					nodes = append(nodes, CompleteNode{Type: RuleOp, Content: top})
				}
				opStack = opStack.push(r)
			} else if expr[st] == '(' {
				opStack = opStack.push(r)
			} else if expr[st] == ')' {
				// 弹出栈中所有运算符直到{
				for top := opStack.peek(); top != 0; top = opStack.peek() {
					opStack, _ = opStack.pop()
					if top != '(' {
						nodes = append(nodes, CompleteNode{Type: RuleOp, Content: top})
					} else {
						break
					}
				}
			} else {
				return &nodes, errors.New("请检查语法")
			}
		}
	}
	for top := opStack.peek(); top != 0; top = opStack.peek() {
		opStack, _ = opStack.pop()
		nodes = append(nodes, CompleteNode{Type: RuleOp, Content: top})
	}
	return &nodes, nil
}
