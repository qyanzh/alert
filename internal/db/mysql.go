package db

import (
	"alert/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DbClient *gorm.DB

func init() {
	var err error
	DbClient, err = gorm.Open(mysql.Open(
		"qyanzh:Qy218456@tcp(111.62.122.250:3306)/test-zqy?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	DbClient.AutoMigrate(
		&model.Alert{},
		&model.Index{},
		&model.Rule{},
		&model.Task{},
	)
}
