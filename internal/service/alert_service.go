package service

import "alert/internal/dao"

type AlertService struct {
	alertDao dao.AlertDao
}

func NewAlertService() *AlertService {
	return &AlertService{alertDao: *dao.NewAlertDao()}
}
