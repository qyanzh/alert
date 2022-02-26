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

func (dao *IndexDao) AddIndex(index *model.Index) int64 {
	result := dao.db.Create(&index)
	return result.RowsAffected
}

func (dao *IndexDao) DeleteIndexByID(ID uint) int64 {
	result := dao.db.Delete(&model.Index{}, ID)
	return result.RowsAffected
}

func (dao *IndexDao) UpdateIndex(index *model.Index) int64 {
	result := dao.db.Save(&index)
	return result.RowsAffected
}

func (dao *IndexDao) SelectIndexByID(ID uint) *model.Index {
	index := model.Index{}
	dao.db.First(&index, ID)
	return &index
}

func (dao *IndexDao) SelectIndexByCode(code string) *model.Index {
	index := model.Index{Code: code}
	dao.db.Where(&index).First(&index)
	return &index
}

func (dao *IndexDao) SelectAllIndexes() *[]model.Index {
	var indexes []model.Index
	dao.db.Find(&indexes)
	return &indexes
}
