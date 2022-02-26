/**
 * @author  qyanzh
 * @create  2022/02/26 21:06
 */

package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID        uint `gorm:"primarykey"`
	RoomID    uint
	GoodID    uint
	UserID    uint
	Turnover  float64
	OrderTime time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (o Order) TableName() string {
	return "orders"
}
