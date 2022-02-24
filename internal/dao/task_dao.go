package dao

import (
	"alert/internal/db"
	"gorm.io/gorm"
)

type TaskDao struct {
	db *gorm.DB
}

func NewTaskDao() *TaskDao {
	return &TaskDao{db: db.DbClient}
}
