/**
 * @author  qyanzh
 * @create  2022/02/26 15:55
 */

package indices

import (
	"alert/internal/dao"
	"alert/internal/model"
	"container/list"
	"fmt"
)

type IndexEvaluator struct {
	indexDao *dao.IndexDao
	orderDao *dao.OrderDao
}

func NewIndexEvaluator(indexDao *dao.IndexDao, orderDao *dao.OrderDao) *IndexEvaluator {
	return &IndexEvaluator{indexDao, orderDao}
}

func (e *IndexEvaluator) Eval(index *model.Index, roomID uint) (float64, error) {
	postExpr, err := postExprFromJson(index.Serialized)
	if err != nil {
		return 0, err
	}
	return e.eval(postExpr, roomID, index.TimeRange)
}

// eval 对后缀表达式求值
func (e *IndexEvaluator) eval(indexExpr *postExpr, roomID, timeRange uint) (float64, error) {
	numStack := list.New() // 存放操作数(float64)
	for _, node := range *indexExpr {
		switch node.NodeType {
		case op:
			operand2, err := popOperand(numStack)
			if err != nil {
				return 0, fmt.Errorf("evaluate failed, handling Op %v: %v", op, err)
			}
			operand1, err := popOperand(numStack)
			if err != nil {
				return 0, fmt.Errorf("evaluate failed, handling Op %v: %v", op, err)
			}
			numStack.PushBack(doOperation(operand1, operand2, node.Op))
		case num:
			numStack.PushBack(node.Num)
		case code:
			indexValue, err := e.extractSubIndexValue(node.Code, roomID, timeRange)
			if err != nil {
				return 0, fmt.Errorf("evaluate failed, extracting subIndex(Code=%s): %v", node.Code, err)
			}
			numStack.PushBack(indexValue)
		case raw:
			r, err := e.orderDao.SelectValue(node.Raw, roomID, timeRange)
			if err != nil {
				return 0, err
			}
			numStack.PushBack(r)
		}
	}
	if numStack.Len() != 1 {
		return 0, fmt.Errorf("evaluate failed, stack: %v", numStack)
	}
	value, ok := numStack.Back().Value.(float64)
	if !ok {
		return 0, fmt.Errorf("evaluate failed, unexpected element:%v", numStack.Back().Value)
	}
	return value, nil
}

func popOperand(numStack *list.List) (float64, error) {
	node := numStack.Back()
	if node == nil {
		return 0, fmt.Errorf("no operands")
	}
	operand := node.Value.(float64)
	numStack.Remove(node)
	return operand, nil
}

func doOperation(operand1, operand2 float64, op rune) float64 {
	var result float64
	switch op {
	case '+':
		result = operand1 + operand2
	case '-':
		result = operand1 - operand2
	case '*':
		result = operand1 * operand2
	case '/':
		result = operand1 / operand2
	}
	return result
}

func (e *IndexEvaluator) extractSubIndexValue(indexCode string, roomID, timeRange uint) (float64, error) {
	var result float64
	index, err := e.indexDao.SelectIndexByCode(indexCode)
	if err != nil {
		return 0, err
	}
	if index.Type == model.ITNormal {
		if timeRange == 0 { // 若父指标无时间范围，使用子指标时间范围
			timeRange = index.TimeRange
		}
		result, err = e.orderDao.SelectValue(index.Expr, roomID, timeRange)
		if err != nil {
			return 0, err
		}
	} else if index.Type == model.ITComputational {
		subExpr, err := postExprFromJson(index.Serialized)
		if err != nil {
			return 0, fmt.Errorf("deserialize index json(IndexCode=%s): %v", indexCode, err)
		}
		result, err = e.eval(subExpr, roomID, timeRange)
		if err != nil {
			return 0, fmt.Errorf("evaluating sub index %v: %v", index, err)
		}
	}
	return result, nil
}
