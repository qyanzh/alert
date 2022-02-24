package dao

import (
	"alert/internal/db"
	"gorm.io/gorm"
)

type RuleDao struct {
	db *gorm.DB
}

func NewRuleDao() *RuleDao {
	return &RuleDao{db: db.DbClient}
}
