/**
 * @author  qyanzh
 * @create  2022/02/26 15:40
 */

package model

import (
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

type IndexType uint8

const (
	Normal IndexType = iota
	Computational
)
