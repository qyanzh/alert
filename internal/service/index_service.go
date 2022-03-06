/**
 * @author  qyanzh
 * @create  2022/03/07 21:53
 */

package service

import (
	"alert/internal/dao"
	"alert/internal/evaluator/indices"
	"alert/internal/model"
	"fmt"
)

type IndexService struct {
	indexDao       *dao.IndexDao
	indexEvaluator *indices.IndexEvaluator
}

func NewIndexService() *IndexService {
	indexDao := dao.NewIndexDao()
	orderDao := dao.NewOrderDao()
	return &IndexService{indexDao: indexDao,
		indexEvaluator: indices.NewIndexEvaluator(indexDao, orderDao)}
}

func (is *IndexService) SelectIndexByID(id uint) (*model.Index, error) {
	return is.indexDao.SelectIndexByID(id)
}

func (is *IndexService) SelectIndexByCode(code string) (*model.Index, error) {
	return is.indexDao.SelectIndexByCode(code)
}

func (is *IndexService) SelectAllIndices() (*[]model.Index, error) {
	return is.indexDao.SelectAllIndices()
}

func (is *IndexService) DeleteIndex(code string) error {
	_, err := is.indexDao.DeleteIndexByCode(code, false)
	return err
}

func (is *IndexService) UpdateIndex(index *model.Index) error {
	_, err := is.indexDao.UpdateIndex(index)
	return err
}

func (is *IndexService) AddIndex(name, code, expr string, timeRange uint) (*model.Index, error) {
	serialized, indexType, err := indices.InfixToPostExprJson(expr)
	if err != nil {
		return nil, fmt.Errorf("adding index(code=%s, expr=%s): %v", code, expr, err)
	}
	newIndex := model.Index{
		Code:       code,
		Name:       name,
		Type:       indexType,
		Expr:       expr,
		Serialized: serialized,
		TimeRange:  timeRange,
	}
	_, err = is.indexDao.AddIndex(&newIndex)
	if err != nil {
		return nil, err
	}
	return &newIndex, nil
}

func (is *IndexService) SelectRoomIndices(codes []string, roomID uint) (map[string]float64, error) {
	indicesBatch, err := is.indexDao.SelectIndexByCodeBatch(codes)
	if err != nil {
		return nil, err
	}
	type Result struct {
		code  string
		value float64
	}
	indexValues := make(chan Result)
	for _, index := range *indicesBatch {
		go func(i model.Index) {
			value, innerError := is.indexEvaluator.Eval(&i, roomID)
			if innerError != nil {
				err = innerError
			}
			indexValues <- Result{code: i.Code, value: value}
		}(index)
	}
	m := make(map[string]float64)
	for range *indicesBatch {
		result := <-indexValues
		m[result.code] = result.value
	}
	return m, err
}
