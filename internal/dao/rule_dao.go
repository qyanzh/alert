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

func (rd *RuleDao) AddRule(rule *model.Rule) (int64, error) {
	var temp int64
	if rd.db.Model(rule).Unscoped().Where("code=?", rule.Code).Count(&temp); temp > 0 {
		var tempRule model.Rule
		rd.db.Model(rule).Unscoped().Where("code=?", rule.Code).First(&tempRule)
		rule.Id = tempRule.Id
		rule.CreatedAt = time.Now()
		rd.db.Model(rule).Unscoped().Where("code=?", rule.Code).Update("deleted_at", nil)
		return rd.UpdateRule(rule)
	}
	result := rd.db.Create(&rule)
	return result.RowsAffected, result.Error
}

func (rd *RuleDao) DeleteRuleByID(ID uint) (int64, error) {
	result := rd.db.Delete(&model.Rule{}, ID)
	return result.RowsAffected, result.Error
}

func (rd *RuleDao) DeleteRuleByCode(code string) (int64, error) {
	result := rd.db.Where("code=?", code).Delete(&model.Rule{})
	return result.RowsAffected, result.Error
}

func (rd *RuleDao) UpdateRule(rule *model.Rule) (int64, error) {
	result := rd.db.Save(&rule)
	return result.RowsAffected, result.Error
}

func (rd *RuleDao) SelectRuleByID(ID uint) (*model.Rule, error) {
	rule := model.Rule{}
	result := rd.db.First(&rule, ID)
	return &rule, result.Error
}

func (rd *RuleDao) SelectRuleByCode(code string) (*model.Rule, error) {
	rule := model.Rule{
		Code: code,
	}
	result := rd.db.Where(&rule).First(&rule)
	return &rule, result.Error
}

func (rd *RuleDao) SelectAllRules() (*[]model.Rule, error) {
	var rules []model.Rule
	result := rd.db.Find(&rules)
	return &rules, result.Error
}
