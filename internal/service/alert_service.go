package service

import (
	"alert/internal/dao"
	"alert/internal/model"
	"time"
)

type AlertService struct {
	alertDao dao.AlertDao
}

func NewAlertService() *AlertService {
	return &AlertService{alertDao: *dao.NewAlertDao()}
}

func (service *AlertService) AddAlert(ruleId uint, t time.Time) error {
	alert := model.Alert{RuleId: ruleId, Time: t}
	_, err := service.alertDao.AddAlert(&alert)
	return err
}

func (service *AlertService) SelectAlert(ruleId uint, startTime time.Time, endTime time.Time) (*[]model.Alert, error) {
	return service.alertDao.SelectAlertByOther(ruleId, startTime, endTime)
}
