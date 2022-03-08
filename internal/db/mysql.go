package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DbClient *gorm.DB

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	var err error
	DbClient, err = gorm.Open(mysql.Open(
		"alert_group_3:alert_group_3@tcp(111.62.122.250:3306)/alert_stc_zqy?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		log.Panic(err)
	}

	//err = DbClient.AutoMigrate(
	//	&model.Alert{},
	//	&model.Index{},
	//	&model.Order{},
	//	&model.Rule{},
	//	&model.Task{},
	//)
	//if err != nil {
	//	log.Panic(err)
	//}
}
