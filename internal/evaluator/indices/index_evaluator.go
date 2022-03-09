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
	var result float64
	var err error
	if index.Type == model.ITNormal {
		result, err = e.evalNormal(index.Expr, roomID, index.TimeRange)
	} else if index.Type == model.ITComputational {
		pe, innerErr := postExprFromJson(index.Serialized)
		if innerErr != nil {
			return 0, innerErr
		}
		result, err = e.evalComputational(pe, roomID, index.TimeRange)
	}
	if err != nil {
		return 0, fmt.Errorf("evaluating expr=%s failed: %v", index.Expr, err)
	}
	return result, nil
}

func (e *IndexEvaluator) evalNormal(expr string, roomID, timeRange uint) (float64, error) {
	r, err := e.orderDao.SelectValue(expr, roomID, timeRange)
	if err != nil {
		return 0, err
	}
	return r, nil
}

// evalComputational 对后缀表达式求值
func (e *IndexEvaluator) evalComputational(indexExpr *postExpr, roomID, timeRange uint) (float64, error) {
	numStack := list.New() // 存放操作数(float64)
	for _, node := range *indexExpr {
		switch node.NodeType {
		case op:
			operand2, err := popOperand(numStack)
			if err != nil {
				return 0, fmt.Errorf("handling Op %v: %v", op, err)
			}
			operand1, err := popOperand(numStack)
			if err != nil {
				return 0, fmt.Errorf("handling Op %v: %v", op, err)
			}
			numStack.PushBack(doOperation(operand1, operand2, node.Op))
		case num:
			numStack.PushBack(node.Num)
		case code:
			indexValue, err := e.extractSubIndexValue(node.Code, roomID, timeRange)
			if err != nil {
				return 0, fmt.Errorf("handling subIndex(Code=%s): %v", node.Code, err)
			}
			numStack.PushBack(indexValue)
		case raw:
			r, err := e.evalNormal(node.Raw, roomID, timeRange)
			if err != nil {
				return 0, fmt.Errorf("handling raw expr=%s: %v", node.Raw, err)
			}
			numStack.PushBack(r)
		}
	}
	if numStack.Len() != 1 {
		return 0, fmt.Errorf("stack: %v", numStack)
	}
	value, ok := numStack.Back().Value.(float64)
	if !ok {
		return 0, fmt.Errorf("unexpected element:%v", numStack.Back().Value)
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
		result, err = e.evalComputational(subExpr, roomID, timeRange)
		if err != nil {
			return 0, fmt.Errorf("evaluating sub index %v: %v", index, err)
		}
	}
	return result, nil
}
