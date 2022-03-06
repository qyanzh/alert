/**
 * @author  qyanzh
 * @create  2022/02/26 22:31
 */

package dao

import (
	"alert/internal/db"
	"alert/internal/model"
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
	return result.RowsAffected, result.Error
}

func (dao *IndexDao) DeleteIndexByID(ID uint) (int64, error) {
	result := dao.db.Delete(&model.Index{}, ID)
	return result.RowsAffected, result.Error
}

func (dao *IndexDao) UpdateIndex(index *model.Index) (int64, error) {
	result := dao.db.Save(index)
	return result.RowsAffected, result.Error
}

func (dao *IndexDao) SelectIndexByID(ID uint) (*model.Index, error) {
	index := model.Index{}
	result := dao.db.First(&index, ID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &index, nil
}

func (dao *IndexDao) SelectIndexByCode(code string) (*model.Index, error) {
	index := model.Index{Code: code}
	result := dao.db.Where(&index).First(&index)
	if result.Error != nil {
		return nil, result.Error
	}
	return &index, nil
}

func (dao *IndexDao) SelectAllIndexes() (*[]model.Index, error) {
	var indexes []model.Index
	result := dao.db.Find(&indexes)
	if result.Error != nil {
		return nil, result.Error
	}
	return &indexes, nil
}
