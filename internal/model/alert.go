package model

import (
	"gorm.io/gorm"
	"time"
)

type Alert struct {
	Id        uint `gorm:"primary_key"`
	Time      time.Time
	RoomId    uint
	RuleId    uint
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
