/**
 * @author  qyanzh
 * @create  2022/02/26 15:56
 */

package test

import (
	"alert/internal/evaluator"
	"fmt"
	"testing"
)

func TestToIndexNodes(t *testing.T) {
	expr := "(index[index1] + index[index2]) * num[2] + num[1] / num[2]"
	nodes := evaluator.ToIndexExpr(expr)
	fmt.Println(nodes)
}

func TestEval(t *testing.T) {
	expr := "n[3.14] * ((n[2] + n[1])+n[3]) / n[2]"
	nodes := evaluator.ToIndexExpr(expr)
	fmt.Println(nodes)
	s := nodes.ToJson()
	fmt.Println(evaluator.IndexExprFromJson(s))
	fmt.Println(nodes.Eval(0, 0))
}

func TestRawExpr(t *testing.T) {
	expr := "raw[sum(turnover) / 2] / index[turnover]"
	nodes := evaluator.ToIndexExpr(expr)
	fmt.Println(nodes)
	fmt.Println(nodes.Eval(0, 120*60))
}
