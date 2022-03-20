package service

import (
	"alert/internal/dao"
	"alert/internal/model"
	"encoding/json"
	"time"
)

type AlertService struct {
	alertDao dao.AlertDao
}

func NewAlertService() *AlertService {
	return &AlertService{alertDao: *dao.NewAlertDao()}
}

func (as *AlertService) AddAlert(ruleId uint, t time.Time, indexNum map[uint]float64) error {
	index, err := json.Marshal(indexNum)
	if err != nil {
		return err
	}
	alert := model.Alert{RuleId: ruleId, Time: t, IndexNum: index}
	_, err = as.alertDao.AddAlert(&alert)
	return err
}

func (as *AlertService) SelectAlert(ruleId uint, startTime time.Time, endTime time.Time) (*[]model.Alert, error) {
	return as.alertDao.SelectAlertByOther(ruleId, startTime, endTime)
}
