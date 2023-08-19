package db

import (
	"log"

	"github.com/slicequeue/go-study-rest-api-board/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/go-board?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("storage err:", err)
	}
	db.Logger.LogMode(logger.Info)
	return db
}

// 모델 등록 부분
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Board{},
		&model.Document{},
	)
}
