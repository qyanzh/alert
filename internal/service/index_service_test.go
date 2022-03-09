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
	allIndices, err := indexService.SelectAllIndices()
	if err != nil {
		t.Error(err)
	}
	ids := make([]uint, len(*allIndices))
	for _, index := range *allIndices {
		ids = append(ids, index.ID)
	}
	results, err := indexService.SelectIndexValuesByIDsAndRoomID(ids, 0)
	if err != nil {
		t.Error(err)
	}
	t.Log(results)

}

func TestAddIndex(t *testing.T) {
	type temp struct {
		code      string
		name      string
		expr      string
		timeRange uint
	}
	exprs := []temp{
		//{"turnover", "总营业额", "sum(turnover)", 0},
		//{"half of turnover recent 20 min", "最近20分钟营业额的一半", "index[turnover] / num[2]", 1200},
		//{"half of turnover recent 1 hour", "最近一小时营业额的一半", "index[turnover] / num[2]", 3600},
		//{"number of orders", "总订单量", "count(*)", 0},
		{"number of orders computational", "总订单量(计算型)", "raw[count(*)]", 0},
	}
	for _, e := range exprs {
		_, err := indexService.AddIndex(e.name, e.code, e.expr, e.timeRange)
		if err != nil {
			t.Error(err)
		}
	}
}
