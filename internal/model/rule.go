package model

import (
	"gorm.io/gorm"
	"time"
)

type Rule struct {
	Id         uint   `gorm:"primary_key"`
	Code       string `gorm:"unique"`
	RoomId     uint
	Name       string
	Type       RuleType
	Expr       string
	Serialized []byte
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type RuleType uint8

const (
	COMPLEXRULE RuleType = iota
	NORMALRULE
)
