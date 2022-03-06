/**
 * @author  qyanzh
 * @create  2022/02/26 15:40
 */

package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Index struct {
	ID         uint   `gorm:"primarykey"`
	Code       string `gorm:"unique"`
	Name       string
	Type       IndexType // 普通型/计算型
	Expr       string    // e.g. (index[index_code] + num[100]) / raw[sum(column)]
	Serialized []byte    // 将Expr解析为IndexExpr后，序列化为json字符串
	TimeRange  uint      // 指标的时间范围，单位：s，0表示不限
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (index *Index) String() string {
	return fmt.Sprintf("Index{ID=%d, Code=%s, Name=%s, Type=%v, Expr=%s, Serialized=%s, timeRange=%d}",
		index.ID, index.Code, index.Name, index.Type, index.Expr, index.Serialized, index.TimeRange)
}

type IndexType uint8

const (
	Normal IndexType = iota
	Computational
)

func (it *IndexType) String() string {
	var s string
	switch *it {
	case Normal:
		s = "普通型"
	case Computational:
		s = "计算型"
	}
	return s
}
