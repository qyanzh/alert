package service

import (
	"testing"
	"time"
)

var alertService AlertService

func init() {
	alertService = *NewAlertService()
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
			err := alertService.AddAlert(value, value2)
			if err != nil {
				print(err.Error())
			}
		}
	}
}
func TestSelectAlert(t *testing.T) {
	alerts, err := alertService.SelectAlert(3, time.Unix(0, 0), time.Unix(0, 0))
	if err != nil {
		print(err.Error())
	}
	for _, value := range *alerts {
		print(value.Id)
	}
	stTime := time.Date(2022, time.March, 2, 15, 30, 30, 0, time.Local)
	edtime := stTime.Add(5 * time.Hour)
	alerts, err = alertService.SelectAlert(3, stTime, edtime)
	if err != nil {
		print(err.Error())
	}
	for _, value := range *alerts {
		print(value.Id)
	}
	alerts, err = alertService.SelectAlert(0, stTime, time.Unix(0, 0))
	if err != nil {
		print(err.Error())
	}
	for _, value := range *alerts {
		print(value.Id)
	}
}
