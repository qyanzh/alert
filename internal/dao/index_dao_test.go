/**
 * @author  qyanzh
 * @create  2022/02/26 22:47
 */

package dao

import (
	"alert/internal/model"
	"testing"
)

var indexDao *IndexDao

func init() {
	indexDao = NewIndexDao()
}

func TestAddIndex(t *testing.T) {
	_, _ = indexDao.DeleteIndexByCode("test_index", true)
	index := model.Index{}
	index.Type = model.ITNormal
	index.Code = "test_index"
	index.Name = "总营业额"
	index.Expr = "sum(turnover)"
	index.TimeRange = 60
	_, err := indexDao.AddIndex(&index)
	if err != nil {
		t.Error(err)
	}
}

func TestAddDuplicateIndex(t *testing.T) {
	index := model.Index{}
	index.Type = model.ITNormal
	index.Code = "test_index"
	index.Name = "总营业额"
	index.Expr = "sum(turnover)"
	index.TimeRange = 60
	_, err := indexDao.AddIndex(&index)
	if err != nil {
		t.Log(err)
	} else {
		t.Error("重复插入相同code的index")
	}
}

func TestDeleteExistIndex(t *testing.T) {
	TestAddIndex(t)
	index, err := indexDao.SelectIndexByCode("test_index")
	if err != nil {
		t.Error(err)
	}
	_, err = indexDao.DeleteIndexByID(index.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateExistIndex(t *testing.T) {
	TestAddIndex(t)
	index, err := indexDao.SelectIndexByCode("test_index")
	if err != nil {
		t.Error(err)
	}
	index.Name = "test_update_" + index.Name
	_, err = indexDao.UpdateIndex(index)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateNonExistIndex(t *testing.T) {
	TestDeleteExistIndex(t)
	index := model.Index{}
	index.Code = "test_index"
	index.Name = "test_update_" + index.Name
	_, err := indexDao.UpdateIndex(&index)
	if err != nil {
		t.Log(err)
	} else {
		t.Error("更新不存在的index")
	}
}

func TestSelectExistIndexByCode(t *testing.T) {
	TestAddIndex(t)
	_, err := indexDao.SelectIndexByCode("test_index")
	if err != nil {
		t.Error(err)
	}
}

func TestSelectNonExistIndexByCode(t *testing.T) {
	TestDeleteExistIndex(t)
	_, err := indexDao.SelectIndexByCode("test_index")
	if err != nil {
		t.Log(err)
	} else {
		t.Error("查出不存在的index")
	}
}

func TestSelectIndexByCodeBatch(t *testing.T) {
	indices, err := indexDao.SelectIndexByCodeBatch([]string{
		"turnover",
		"half of turnover recent 3 min",
		"turnover*4",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(indices)
}
