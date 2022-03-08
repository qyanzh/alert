/**
 * @author  qyanzh
 * @create  2022/03/07 1:54
 */

package service

import (
	"testing"
)

var indexService *IndexService

func init() {
	indexService = NewIndexService()
}

func TestSelectRoomIndicesByCodes(t *testing.T) {
	indexBatch, err := indexService.SelectAllIndices()
	if err != nil {
		t.Error(err)
	}
	codes := make([]string, len(*indexBatch))
	for _, index := range *indexBatch {
		codes = append(codes, index.Code)
	}
	indices, err := indexService.SelectIndexValuesByCodesAndRoomID(codes, 0)
	t.Log(indices)
}

func TestSelectRoomIndicesByIDs(t *testing.T) {
	indexBatch, err := indexService.SelectAllIndices()
	if err != nil {
		t.Error(err)
	}
	ids := make([]uint, len(*indexBatch))
	for _, index := range *indexBatch {
		ids = append(ids, index.ID)
	}
	indices, err := indexService.SelectIndexValuesByIDsAndRoomID(ids, 0)
	t.Log(indices)
}

func TestAddIndex(t *testing.T) {
	_, err := indexService.indexDao.DeleteIndexByCode("turnover_test_service3")
	index, err := indexService.AddIndex("营业额", "turnover_test_service3", "raw[sum(turnover)]", 0)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", index)
}
