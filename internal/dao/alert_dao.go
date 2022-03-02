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

func (dao *AlertDao) AddAlert(alert *model.Rule) (int64, error) {
	result := dao.db.Create(&alert)
	return result.RowsAffected, result.Error
}
func (dao *AlertDao) SelectAlertByID(ID uint) (*model.Alert, error) {
	alert := model.Alert{}
	result := dao.db.First(&alert, ID)
	return &alert, result.Error
}
func (dao *AlertDao) SelectAlertByOther(roomId uint, ruleId uint, startTime time.Time, endTime time.Time) (*[]model.Alert, error) {
	nowDb := dao.db
	if roomId != -1 {
		nowDb = nowDb.Where("room_id=?", roomId)
	}
	if ruleId != -1 {
		nowDb = nowDb.Where("ruleId=?", ruleId)
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
