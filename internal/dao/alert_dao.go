package dao

import (
	"alert/internal/db"
	"alert/internal/model"
	"gorm.io/gorm"
	"time"
)

type AlertDao struct {
	db *gorm.DB
}

func NewAlertDao() *AlertDao {
	return &AlertDao{db: db.DbClient}
}

func (ad *AlertDao) AddAlert(alert *model.Alert) (int64, error) {
	result := ad.db.Create(&alert)
	return result.RowsAffected, result.Error
}
func (ad *AlertDao) SelectAlertByID(ID uint) (*model.Alert, error) {
	alert := model.Alert{}
	result := ad.db.First(&alert, ID)
	return &alert, result.Error
}
func (ad *AlertDao) SelectAlertByOther(ruleId uint, startTime time.Time, endTime time.Time) (*[]model.Alert, error) {
	nowDb := ad.db
	if ruleId != 0 {
		nowDb = nowDb.Where("rule_id=?", ruleId)
	}
	if startTime != time.Unix(0, 0) {
		nowDb = nowDb.Where("time>?", startTime)
	}
	if endTime != time.Unix(0, 0) {
		nowDb = nowDb.Where("time<?", endTime)
	}
	var alerts []model.Alert
	result := nowDb.Find(&alerts)
	return &alerts, result.Error
}
