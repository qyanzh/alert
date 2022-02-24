package dao

import (
	"alert/internal/db"
	"gorm.io/gorm"
)

type IndexDao struct {
	db *gorm.DB
}

func NewIndexDao() *IndexDao {
	return &IndexDao{db: db.DbClient}
}
