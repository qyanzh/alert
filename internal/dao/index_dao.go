/**
 * @author  qyanzh
 * @create  2022/02/26 22:31
 */

package dao

import (
	"alert/internal/db"
	"alert/internal/model"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type IndexDao struct {
	db *gorm.DB
}

func NewIndexDao() *IndexDao {
	return &IndexDao{db: db.DbClient}
}

func (id *IndexDao) AddIndex(index *model.Index) (int64, error) {
	result := id.db.Create(index)
	if result.Error != nil {
		var uniqueError *mysql.MySQLError
		if errors.As(result.Error, &uniqueError) && uniqueError.Number == 1062 {
			var sameCodeIndex model.Index
			innerResult := id.db.Model(index).Unscoped().Where("code = ?", index.Code).First(&sameCodeIndex)
			if innerResult.Error != nil {
				return 0, fmt.Errorf("adding duplicate index %v: %v", *index, innerResult.Error)
			}
			if sameCodeIndex.DeletedAt.Valid {
				id.db.Unscoped().Where("code = ?", index.Code).Delete(&model.Index{})
				return id.AddIndex(index)
			}
		}
		return 0, fmt.Errorf("adding index %v: %v", *index, result.Error)
	}
	return result.RowsAffected, nil
}

func (id *IndexDao) DeleteIndexByID(ID uint) (int64, error) {
	result := id.db.Delete(&model.Index{}, ID)
	if result.Error != nil {
		return 0, fmt.Errorf("deleting index by id=%d: %v", ID, result.Error)
	}
	return result.RowsAffected, nil
}

func (id *IndexDao) DeleteIndexByCode(code string) (int64, error) {
	result := id.db.Where("code = ?", code).Delete(&model.Index{})
	if result.Error != nil {
		return 0, fmt.Errorf("deleting index by code=%s: %v", code, result.Error)
	}
	return result.RowsAffected, nil
}

func (id *IndexDao) UpdateIndex(index *model.Index) (int64, error) {
	result := id.db.Save(index)
	if result.Error != nil {
		return 0, fmt.Errorf("updating index %v: %v", *index, result.Error)
	}
	return result.RowsAffected, nil
}

func (id *IndexDao) SelectIndexByID(ID uint) (*model.Index, error) {
	index := model.Index{}
	result := id.db.First(&index, ID)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting index by id=%d: %v", ID, result.Error)
	}
	return &index, nil
}

func (id *IndexDao) SelectIndexByIDs(ids []uint) (*[]model.Index, error) {
	indices := make([]model.Index, len(ids))
	result := id.db.Where("id IN ?", ids).Find(&indices)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting index by ids=%s: %v", ids, result.Error)
	}
	return &indices, nil
}

func (id *IndexDao) SelectIndexByCode(code string) (*model.Index, error) {
	index := model.Index{Code: code}
	result := id.db.Where(&index).First(&index)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting index by code=%s: %v", code, result.Error)
	}
	return &index, nil
}

func (id *IndexDao) SelectIndexByCodes(codes []string) (*[]model.Index, error) {
	indices := make([]model.Index, len(codes))
	result := id.db.Where("code IN ?", codes).Find(&indices)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting index by codes=%s: %v", codes, result.Error)
	}
	return &indices, nil
}

func (id *IndexDao) SelectAllIndices() (*[]model.Index, error) {
	var indices []model.Index
	result := id.db.Find(&indices)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting all indices: %v", result.Error)
	}
	return &indices, nil
}
