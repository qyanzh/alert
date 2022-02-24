package dao

import (
	"alert/internal/db"
	"gorm.io/gorm"
)

type AlertDao struct {
	db *gorm.DB
}

func NewAlertDao() *AlertDao {
	return &AlertDao{db: db.DbClient}
}
