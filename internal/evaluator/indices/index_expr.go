/**
 * @author  qyanzh
 * @create  2022/03/06 21:33
 */

package indices

import (
	"container/list"
	"encoding/json"
	"fmt"
	"strconv"
	"unicode/utf8"
)

// postExpr 指标后序表达式
type postExpr []*postNode

// postNode 后序表达式中的节点
type postNode struct {
	NodeType nodeType
	Op       rune
	Num      float64
	Code     string
	Raw      string
}

type nodeType uint8

const (
	op   nodeType = iota // 运算符
	num                  // 数字
	code                 // 子指标，需要查表得到sql表达式
	raw                  // sql表达式，直接拼接在select中
)

func InfixToPostExprJson(infix string) ([]byte, error) {
	postExpr, err := infixToPostExpr(infix)
	if err != nil {
		return nil, err
	}
	js, err := postExpr.json()
	if err != nil {
		return nil, fmt.Errorf("post expr(from %s) to json: %v", infix, err)
	}
	return js, nil
}

func postExprFromJson(bs []byte) (*postExpr, error) {
	var expr postExpr
	err := json.Unmarshal(bs, &expr)
	return &expr, err
}

func (ie *postExpr) json() ([]byte, error) {
	return json.Marshal(ie)
}

// infixToPostExpr 解析中缀表达式为后缀表达式结点数组
func infixToPostExpr(infix string) (*postExpr, error) {
	var pe postExpr
	opStack := list.New() // 存放运算符和括号(rune)
	for i := 0; i < len(infix); {
		r, size := utf8.DecodeRuneInString(infix[i:])
		if r == ' ' {
			i += size
			continue
		}
		if isNodePrefix(r) {
			node, offset, err := nextNode(r, infix[i:])
			if err != nil {
				return nil, err
			}
			pe = append(pe, node)
			i += offset
		} else {
			if isOp(r) {
				// 弹出栈中优先级>=当前运算符的运算符
				for top := opStack.Back(); top != nil; top = opStack.Back() {
					topv := top.Value.(rune)
					if topv != '(' && opGE(topv, r) {
						pe = append(pe, &postNode{NodeType: op, Op: topv})
						opStack.Remove(top)
					} else {
						break
					}
				}
				opStack.PushBack(r)
			} else if r == '(' {
				opStack.PushBack(r)
			} else if r == ')' {
				// 弹出栈中所有运算符直到(
				for top := opStack.Back(); top != nil; top = opStack.Back() {
					topv := top.Value.(rune)
					opStack.Remove(top)
					if topv != '(' {
						pe = append(pe, &postNode{NodeType: op, Op: topv})
					} else {
						break
					}
				}
			} else {
				return nil, fmt.Errorf("syntax error: unexpected rune %c", r)
			}
			i += size
		}
	}
	for top := opStack.Back(); top != nil; top = opStack.Back() {
		topv := top.Value.(rune)
		if !isOp(topv) {
			return nil, fmt.Errorf("illegal opStack(%v) while converting infixExpr(%s)", opStack, infix)
		}
		pe = append(pe, &postNode{NodeType: op, Op: top.Value.(rune)})
		opStack.Remove(top)
	}
	return &pe, nil
}

func isNodePrefix(r rune) bool {
	return r == 'i' || r == 'n' || r == 'r'
}

func nextNode(r rune, s string) (*postNode, int, error) {
	content, offset := decodeContentBetween(s, '[', ']')
	var pn postNode
	switch r {
	case 'i':
		pn.NodeType = code
		pn.Code = content
	case 'n':
		val, err := strconv.ParseFloat(content, 64)
		if err != nil {
			return nil, 0, fmt.Errorf("handling expr node content(%s): %v", content, err)
		}
		pn.NodeType = num
		pn.Num = val
	case 'r':
		pn.NodeType = raw
		pn.Raw = content
	}
	return &pn, offset, nil
}

// 取出下一个beginRune~endRune中的内容
func decodeContentBetween(expr string, beginRune, endRune rune) (content string, offset int) {
	var begin, end int
	for i := 0; i < len(expr); {
		r, size := utf8.DecodeRuneInString(expr[i:])
		if begin == 0 && r == beginRune {
			begin = i + size
		} else if r == endRune {
			end = i
			content = expr[begin:end]
			offset = i + size
			break
		}
		i += size
	}
	return
}

// 运算符优先级
var opPriority = map[rune]int8{
	'+': 0,
	'-': 0,
	'*': 1,
	'/': 1,
}

func isOp(s rune) bool {
	return s == '+' || s == '-' || s == '*' || s == '/'
}

func opGE(o1, o2 rune) bool {
	return opPriority[o1]-opPriority[o2] >= 0
}

func (ie *postExpr) String() string {
	ret := "postExpr:"
	for _, node := range *ie {
		var i string
		switch node.NodeType {
		case op:
			i = "op$" + string(node.Op) + "$"
		case num:
			i = "num$" + fmt.Sprintf("%f", node.Num) + "$"
		case code:
			i = "code$" + node.Code + "$"
		case raw:
			i = "raw$" + node.Raw + "$"
		}
		ret = fmt.Sprintf("%s %s", ret, i)
	}
	return ret
}
