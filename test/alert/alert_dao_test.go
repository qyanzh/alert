package alert

import (
	"alert/internal/dao"
	"alert/internal/model"
	"testing"
	"time"
)

var alertDao dao.AlertDao

func init() {
	alertDao = *dao.NewAlertDao()
}

func TestAddAlert(t *testing.T) {
	ruleIds := []uint{1, 2, 3}
	now := time.Now()
	times := []time.Time{
		now.Add(time.Hour),
		now.Add(5 * time.Hour),
		now.Add(10 * time.Hour),
	}
	for _, value := range ruleIds {
		for _, value2 := range times {
			alert := model.Alert{}
			alert.Time = value2
			alert.RuleId = value
			_, err := alertDao.AddAlert(&alert)
			if err != nil {
				print(err.Error())
			}
		}
	}
}
func TestSelectAlert(t *testing.T) {
	alerts, err := alertDao.SelectAlertByOther(3, time.Unix(0, 0), time.Unix(0, 0))
	if err != nil {
		print(err.Error())
	}
	for _, value := range *alerts {
		print(value.Id)
	}
	stTime := time.Date(2022, time.March, 2, 15, 30, 30, 0, time.Local)
	edtime := stTime.Add(5 * time.Hour)
	alerts, err = alertDao.SelectAlertByOther(3, stTime, edtime)
	if err != nil {
		print(err.Error())
	}
	for _, value := range *alerts {
		print(value.Id)
	}
	alerts, err = alertDao.SelectAlertByOther(0, stTime, time.Unix(0, 0))
	if err != nil {
		print(err.Error())
	}
	for _, value := range *alerts {
		print(value.Id)
	}
}
