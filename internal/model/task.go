/**
 * @author  qyanzh
 * @create  2022/03/07 17:38
 */

package model

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID        uint   `gorm:"primarykey"`
	Code      string `gorm:"unique"`
	Name      string
	RuleID    uint
	Frequency uint // 任务频率，单位:s
	Enable    bool
	Status    TaskStatus
	NextTime  time.Time // 运行条件：Enable && Status == TSReady && time.Now().After(NextTime)
	Msg       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type TaskStatus uint8

const (
	TSReady TaskStatus = iota + 1
	TSRunning
	TSFail
)
