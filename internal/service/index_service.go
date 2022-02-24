package service

import "alert/internal/dao"

type IndexService struct {
	alertDao dao.IndexDao
}

func NewIndexService() *IndexService {
	return &IndexService{alertDao: *dao.NewIndexDao()}
}
