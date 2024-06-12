package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	dbInstance *gorm.DB
)

func Mysql() *gorm.DB {
	if dbInstance == nil {
		dsn := "root:12345678@tcp(127.0.0.1:3306)/av3?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}
		dbInstance = db
	}
	return dbInstance
}
