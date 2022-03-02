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
	return &RuleDao{db: db.DbClient}
}

func (dao *RuleDao) AddRule(rule *model.Rule) (int64, error) {
	var temp int64
	if dao.db.Unscoped().Where("code=?", rule.Code).Count(&temp); temp > 0 {
		result := dao.db.Omit("id").Save(rule).
			Updates(map[string]interface{}{"created_at": time.Now(), "updated_at": time.Now(), "deleted_at": nil})
		return result.RowsAffected, result.Error
	}
	result := dao.db.Create(&rule)
	return result.RowsAffected, result.Error
}

func (dao *RuleDao) DeleteRuleByID(ID uint) (int64, error) {
	result := dao.db.Delete(&model.Rule{}, ID)
	return result.RowsAffected, result.Error
}

func (dao *RuleDao) UpdateRule(rule *model.Rule) (int64, error) {
	result := dao.db.Omit("id").Omit("created_at").Save(&rule).Update("updated_at", time.Now())
	return result.RowsAffected, result.Error
}

func (dao *RuleDao) SelectRuleByID(ID uint) (*model.Rule, error) {
	rule := model.Rule{}
	result := dao.db.First(&rule, ID)
	return &rule, result.Error
}

func (dao *RuleDao) SelectRuleByCode(code string) (*model.Rule, error) {
	rule := model.Rule{Code: code}
	result := dao.db.Where(&rule).First(&rule)
	return &rule, result.Error
}

func (dao *RuleDao) SelectAllRules() (*[]model.Rule, error) {
	var rules []model.Rule
	result := dao.db.Find(&rules)
	return &rules, result.Error
}
