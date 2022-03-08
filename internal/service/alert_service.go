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

func (as *AlertService) AddAlert(ruleId uint, t time.Time) error {
	alert := model.Alert{RuleId: ruleId, Time: t}
	_, err := as.alertDao.AddAlert(&alert)
	return err
}

func (as *AlertService) SelectAlert(ruleId uint, startTime time.Time, endTime time.Time) (*[]model.Alert, error) {
	return as.alertDao.SelectAlertByOther(ruleId, startTime, endTime)
}
