/**
 * @author  qyanzh
 * @create  2022/02/26 22:31
 */

package dao

import (
	"alert/internal/db"
	"alert/internal/model"
	"fmt"
	"gorm.io/gorm"
)

type IndexDao struct {
	db *gorm.DB
}

func NewIndexDao() *IndexDao {
	return &IndexDao{db: db.DbClient}
}

func (dao *IndexDao) AddIndex(index *model.Index) (int64, error) {
	result := dao.db.Create(index)
	if result.Error != nil {
		return 0, fmt.Errorf("adding index %v: %v", *index, result.Error)
	}
	return result.RowsAffected, nil
}

func (dao *IndexDao) DeleteIndexByID(ID uint) (int64, error) {
	result := dao.db.Delete(&model.Index{}, ID)
	if result.Error != nil {
		return 0, fmt.Errorf("deleting index by id=%d: %v", ID, result.Error)
	}
	return result.RowsAffected, nil
}

func (dao *IndexDao) DeleteIndexByCode(code string, permanent bool) (int64, error) {
	where := dao.db.Where("code = ?", code)
	if permanent {
		where = where.Unscoped()
	}
	result := where.Delete(&model.Index{})
	if result.Error != nil {
		return 0, fmt.Errorf("deleting index by code=%s: %v", code, result.Error)
	}
	return result.RowsAffected, nil
}

func (dao *IndexDao) UpdateIndex(index *model.Index) (int64, error) {
	result := dao.db.Save(index)
	if result.Error != nil {
		return 0, fmt.Errorf("updating index %v: %v", *index, result.Error)
	}
	return result.RowsAffected, nil
}

func (dao *IndexDao) SelectIndexByID(ID uint) (*model.Index, error) {
	index := model.Index{}
	result := dao.db.First(&index, ID)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting index by id=%d: %v", ID, result.Error)
	}
	return &index, nil
}

func (dao *IndexDao) SelectIndexByCode(code string) (*model.Index, error) {
	index := model.Index{Code: code}
	result := dao.db.Where(&index).First(&index)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting index by code=%s: %v", code, result.Error)
	}
	return &index, nil
}

func (dao *IndexDao) SelectAllIndices() (*[]model.Index, error) {
	var indices []model.Index
	result := dao.db.Find(&indices)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting all indices: %v", result.Error)
	}
	return &indices, nil
}
