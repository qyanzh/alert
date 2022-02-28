package dao

import (
	"alert/internal/db"
	"alert/internal/model"
	"gorm.io/gorm"
)

type RuleDao struct {
	db *gorm.DB
}

func NewRuleDao() *RuleDao {
	return &RuleDao{db: db.DbClient}
}

func (dao *RuleDao) AddRule(rule *model.Rule) int64 {
	result := dao.db.Create(&rule)
	return result.RowsAffected
}

func (dao *RuleDao) DeleteRuleByID(ID uint) int64 {
	result := dao.db.Delete(&model.Rule{}, ID)
	return result.RowsAffected
}

func (dao *RuleDao) UpdateRule(rule *model.Rule) int64 {
	result := dao.db.Save(&rule)
	return result.RowsAffected
}

func (dao *RuleDao) SelectRuleByID(ID uint) *model.Rule {
	rule := model.Rule{}
	dao.db.First(&rule, ID)
	return &rule
}

func (dao *RuleDao) SelectRuleByCode(code string) *model.Rule {
	rule := model.Rule{Code: code}
	dao.db.Where(&rule).First(&rule)
	return &rule
}

func (dao *RuleDao) SelectAllRules() *[]model.Index {
	var rules []model.Index
	dao.db.Find(&rules)
	return &rules
}
