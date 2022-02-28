/**
 * @author  qyanzh
 * @create  2022/02/26 22:47
 */

package index

import (
	"alert/internal/dao"
	"alert/internal/evaluator"
	"alert/internal/model"
	"fmt"
	"testing"
)

var indexDao dao.IndexDao

func init() {
	indexDao = *dao.NewIndexDao()
}

func TestAddNormalIndex(t *testing.T) {
	index := model.Index{}
	index.Type = model.Normal
	index.Code = "turnover"
	index.Name = "总营业额"
	index.Expr = "turnover"
	indexDao.AddIndex(&index)
}

func TestAddComputationalIndex(t *testing.T) {
	index := model.Index{}
	index.Type = model.Computational
	index.Code = "half of turnover recent 1 hour"
	index.Name = "最近一小时营业额的一半"
	index.Expr = "index[turnover] / num[2]"
	indexExpr := evaluator.ToIndexExpr(index.Expr)
	fmt.Println(indexExpr)
	index.Serialized = indexExpr.ToJson()
	index.TimeRange = 60 * 60
	indexDao.AddIndex(&index)
}

func TestAddComputationalIndex2(t *testing.T) {
	index := model.Index{}
	index.Type = model.Computational
	index.Code = "half of turnover recent 20 min"
	index.Name = "最近20分钟营业额的一半"
	index.Expr = "index[turnover] / num[2]"
	indexExpr := evaluator.ToIndexExpr(index.Expr)
	fmt.Println(indexExpr)
	index.Serialized = indexExpr.ToJson()
	index.TimeRange = 20 * 60
	indexDao.AddIndex(&index)
}

func TestComputationalIndex2(t *testing.T) {
	code := "half of turnover recent 20 min"
	index := indexDao.SelectIndexByCode(code)
	fmt.Printf("%+v\n", index)
	result := evaluator.IndexExprFromJson(index.Serialized).Eval(0, index.TimeRange)
	fmt.Printf("%s: %f\n", index.Name, result)
}

func TestAddComputationalIndex3(t *testing.T) {
	index := model.Index{}
	index.Type = model.Computational
	index.Code = "half of turnover recent 3 min"
	index.Name = "最近3分钟营业额的一半"
	index.Expr = "index[turnover] / num[2]"
	indexExpr := evaluator.ToIndexExpr(index.Expr)
	fmt.Println(indexExpr)
	index.Serialized = indexExpr.ToJson()
	index.TimeRange = 3 * 60
	indexDao.AddIndex(&index)
}

func TestComputationalIndex3(t *testing.T) {
	code := "half of turnover recent 3 min"
	index := indexDao.SelectIndexByCode(code)
	fmt.Printf("%+v\n", index)
	result := evaluator.IndexExprFromJson(index.Serialized).Eval(0, index.TimeRange)
	fmt.Printf("%s: %f\n", index.Name, result)
}
