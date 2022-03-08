package dao

import (
	"alert/internal/db"
	"alert/internal/model"
	"gorm.io/gorm"
	"time"
)

type RuleDao struct {
	db *gorm.DB
}

func NewRuleDao() *RuleDao {
	return &RuleDao{
		db: db.DbClient,
	}
}

func (dao *RuleDao) AddRule(rule *model.Rule) (int64, error) {
	var temp int64
	if dao.db.Model(rule).Unscoped().Where("code=?", rule.Code).Count(&temp); temp > 0 {
		var tempRule model.Rule
		dao.db.Model(rule).Unscoped().Where("code=?", rule.Code).First(&tempRule)
		rule.Id = tempRule.Id
		rule.CreatedAt = time.Now()
		dao.db.Model(rule).Unscoped().Where("code=?", rule.Code).Update("deleted_at", nil)
		return dao.UpdateRule(rule)
	}
	result := dao.db.Create(&rule)
	return result.RowsAffected, result.Error
}

func (dao *RuleDao) DeleteRuleByID(ID uint) (int64, error) {
	result := dao.db.Delete(&model.Rule{}, ID)
	return result.RowsAffected, result.Error
}

func (dao *RuleDao) DeleteRuleByCode(code string) (int64, error) {
	result := dao.db.Where("code=?", code).Delete(&model.Rule{})
	return result.RowsAffected, result.Error
}

func (dao *RuleDao) UpdateRule(rule *model.Rule) (int64, error) {
	result := dao.db.Save(&rule)
	return result.RowsAffected, result.Error
}

func (dao *RuleDao) SelectRuleByID(ID uint) (*model.Rule, error) {
	rule := model.Rule{}
	result := dao.db.First(&rule, ID)
	return &rule, result.Error
}

func (dao *RuleDao) SelectRuleByCode(code string) (*model.Rule, error) {
	rule := model.Rule{
		Code: code,
	}
	result := dao.db.Where(&rule).First(&rule)
	return &rule, result.Error
}

func (dao *RuleDao) SelectAllRules() (*[]model.Rule, error) {
	var rules []model.Rule
	result := dao.db.Find(&rules)
	return &rules, result.Error
}
