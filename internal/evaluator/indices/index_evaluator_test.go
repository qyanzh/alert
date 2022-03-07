/**
 * @author  qyanzh
 * @create  2022/02/26 15:56
 */

package indices

import (
	"alert/internal/dao"
	"alert/internal/model"
	"testing"
)

var indexEvaluator *IndexEvaluator

func init() {
	indexEvaluator = NewIndexEvaluator(dao.NewIndexDao(), dao.NewOrderDao())
}

func TestToPostExpr(t *testing.T) {
	expr := "(index[index1] + index[index2]) * num[2] + num[1] / num[2]"
	pe, err := infixToPostExpr(expr)
	if err != nil {
		t.Error(err)
	}
	t.Log(pe)
}

func TestEval(t *testing.T) {
	expr := "n[3.14] * ((n[2] + n[1]) + n[3]) / n[2]"
	postExpr, err := infixToPostExpr(expr)
	if err != nil {
		t.Error(err)
	}
	result, err := indexEvaluator.eval(postExpr, 0, 0)
	t.Log("expr: " + expr)
	t.Log(postExpr)
	t.Log(result)
}

func TestRawExpr(t *testing.T) {
	expr := "raw[sum(turnover)/2] / index[turnover]"
	postExpr, err := infixToPostExpr(expr)
	if err != nil {
		t.Error(err)
	}
	t.Log(postExpr)
	result, err := indexEvaluator.eval(postExpr, 0, 0)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestComputationalIndex(t *testing.T) {
	code := "turnover*4"
	_, err := indexEvaluator.indexDao.DeleteIndexByCode(code, true)
	if err != nil {
		t.Error(err)
	}
	index := model.Index{}
	index.Code = code
	index.Type = model.ITComputational
	index.Expr = "index[turnover*2] * num[2]"
	json, _, err := InfixToPostExprJson(index.Expr)
	if err != nil {
		t.Error(err)
	}
	index.Serialized = json
	_, _ = indexEvaluator.indexDao.UpdateIndex(&index)

	result, err := indexEvaluator.Eval(&index, 0)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSyntaxErrorExprNoNum(t *testing.T) {
	index := model.Index{}
	index.Expr = "index[turnover*2] * 2"
	_, _, err := InfixToPostExprJson(index.Expr)
	if err != nil {
		t.Log(err)
	}
}

func TestSyntaxErrorExprCapitalNum(t *testing.T) {
	index := model.Index{}
	index.Expr = "index[turnover*2} * num[2]"
	pe, err := infixToPostExpr(index.Expr)
	t.Log(pe)
	json, _, err := InfixToPostExprJson(index.Expr)
	if err != nil {
		t.Fatal(err)
	}
	index.Serialized = json
	result, err := indexEvaluator.Eval(&index, 0)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
