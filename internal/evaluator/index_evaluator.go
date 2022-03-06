/**
 * @author  qyanzh
 * @create  2022/02/26 15:55
 */

package evaluator

import (
	"alert/internal/dao"
	"alert/internal/model"
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"unicode/utf8"
)

type IndexNodeType uint8

const (
	op        IndexNodeType = iota // 运算符
	num                            // 数字
	indexCode                      // 子指标，需要查表得到sql表达式
	rawExpr                        // sql表达式，直接拼接在select中
)

// IndexNode 后序表达式中的节点
type IndexNode struct {
	IndexNodeType IndexNodeType
	Op            rune
	Num           float64
	IndexCode     string
	RawExpr       string
}

// IndexExpr 后序表达式
type IndexExpr []IndexNode

func IndexExprFromJson(bs []byte) *IndexExpr {
	var expr IndexExpr
	err := json.Unmarshal(bs, &expr)
	if err != nil {
		log.Panicln(err)
	}
	return &expr
}

func (ie *IndexExpr) ToJson() []byte {
	s, err := json.Marshal(ie)
	if err != nil {
		log.Panicln(err)
	}
	return s
}

func (ie *IndexExpr) String() string {
	var ret string
	for _, node := range *ie {
		var i string
		switch node.IndexNodeType {
		case op:
			i = string(node.Op)
		case num:
			i = fmt.Sprintf("%f", node.Num)
		case indexCode:
			i = node.IndexCode
		case rawExpr:
			i = node.RawExpr
		}
		ret = fmt.Sprintf("%s %s", ret, i)
	}
	return ret
}

// Eval 对后缀表达式求值
func (ie IndexExpr) Eval(roomID uint, timeRange uint) float64 {
	numStack := list.New() // 存放操作数(float64)
	for _, node := range ie {
		switch node.IndexNodeType {
		case op:
			r := numStack.Remove(numStack.Back()).(float64)
			l := numStack.Remove(numStack.Back()).(float64)
			switch node.Op {
			case '+':
				numStack.PushBack(l + r)
			case '-':
				numStack.PushBack(l - r)
			case '*':
				numStack.PushBack(l * r)
			case '/':
				numStack.PushBack(l / r)
			}
		case num:
			numStack.PushBack(node.Num)
		case indexCode:
			// TODO dao成员
			indexDao := dao.NewIndexDao()
			index, err := indexDao.SelectIndexByCode(node.IndexCode)
			if err != nil {
				log.Panicln(err)
			}
			orderDao := dao.NewOrderDao()
			if index.Type == model.Normal {
				if timeRange == 0 { // 若父指标无时间范围，使用子指标时间范围
					timeRange = index.TimeRange
				}
				r, err := orderDao.SelectValue(index.Expr, roomID, timeRange)
				if err != nil {
					log.Panicln(err)
				}
				numStack.PushBack(r)
			} else if index.Type == model.Computational {
				// TODO 测试
				subExpr := IndexExprFromJson(index.Serialized)
				r := subExpr.Eval(roomID, timeRange)
				numStack.PushBack(r)
			}
		case rawExpr:
			orderDao := dao.NewOrderDao()
			r, err := orderDao.SelectValue(node.RawExpr, roomID, timeRange)
			if err != nil {
				log.Panicln(err)
			}
			numStack.PushBack(r)
		}
	}
	return numStack.Front().Value.(float64)
}

// ToIndexExpr 解析中缀表达式为后缀表达式结点数组
func ToIndexExpr(expr string) *IndexExpr {
	nodes := make(IndexExpr, 0)
	opStack := list.New()
	for i := 0; i < len(expr); {
		r, size := utf8.DecodeRuneInString(expr[i:])
		if r == 'i' || r == 'n' || r == 'r' {
			content, offset := decodeContent(expr[i:])
			if r == 'i' {
				nodes = append(nodes, IndexNode{IndexNodeType: indexCode, IndexCode: content})
			} else if r == 'n' {
				val, _ := strconv.ParseFloat(content, 64)
				nodes = append(nodes, IndexNode{IndexNodeType: num, Num: val})
			} else if r == 'r' {
				nodes = append(nodes, IndexNode{IndexNodeType: rawExpr, RawExpr: content})
			}
			i += offset
		} else {
			if isOp(r) {
				// 弹出栈中优先级>=当前运算符的运算符
				for top := opStack.Back(); top != nil; top = opStack.Back() {
					topv := top.Value.(rune)
					if topv != '(' && opGE(topv, r) {
						nodes = append(nodes, IndexNode{IndexNodeType: op, Op: topv})
						opStack.Remove(top)
					} else {
						break
					}
				}
				opStack.PushBack(r)
			} else if r == '(' {
				opStack.PushBack(r)
			} else if r == ')' {
				// 弹出栈中所有运算符直到[
				for top := opStack.Back(); top != nil; top = opStack.Back() {
					topv := top.Value.(rune)
					opStack.Remove(top)
					if topv != '(' {
						nodes = append(nodes, IndexNode{IndexNodeType: op, Op: topv})
					} else {
						break
					}
				}
			}
			i += size
		}
	}
	for top := opStack.Back(); top != nil; top = opStack.Back() {
		nodes = append(nodes, IndexNode{IndexNodeType: op, Op: top.Value.(rune)})
		opStack.Remove(top)
	}
	return &nodes
}

// 取出xxx[...]中的内容
func decodeContent(expr string) (content string, offset int) {
	var begin, end int
	for i := 0; i < len(expr); {
		r, size := utf8.DecodeRuneInString(expr[i:])
		if r == '[' {
			begin = i + size
		} else if r == ']' {
			end = i
			content = expr[begin:end]
			offset = i + size
			break
		}
		i += size
	}
	return
}

func isOp(s rune) bool {
	return s == '+' || s == '-' || s == '*' || s == '/'
}

// 运算符优先级
var opPriority = map[rune]int8{
	'+': 0,
	'-': 0,
	'*': 1,
	'/': 1,
}

func opGE(o1, o2 rune) bool {
	return opPriority[o1]-opPriority[o2] >= 0
}

type stack []rune

func (s stack) push(v rune) stack {
	return append(s, v)
}

func (s stack) pop() (stack, rune) {
	l := len(s)
	if l == 0 {
		return s[:], 0
	}
	return s[:l-1], s[l-1]
}

func (s stack) peek() rune {
	l := len(s)
	if l == 0 {
		return 0
	}
	return s[len(s)-1]
}
