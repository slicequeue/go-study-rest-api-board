package db

import (
	"log"
	"time"

	"github.com/slicequeue/go-study-rest-api-board/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/go-board?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB,_ := db.DB()
	sqlDB.SetMaxIdleConns(2) // 10
	sqlDB.SetMaxOpenConns(5) // 100
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		log.Fatalln("storage err:", err)
	}
	db.Logger.LogMode(logger.Info)
	return db
}

var testDB *gorm.DB

func TestDB() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/go-board-test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	testDB = db
	if err != nil {
		log.Fatalln("storage err:", err)
	}
	return db
}

func DropTestDB() error {
	testDB.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	testDB.Exec("SET @tables = NULL")
	testDB.Exec("SELECT GROUP_CONCAT(table_name) INTO @tables FROM information_schema.tables WHERE table_schema = (SELECT DATABASE())")
	testDB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	testDB.Exec("SET @drop_query = CONCAT('DROP TABLE IF EXISTS ', @tables)")
	testDB.Exec("PREPARE stmt FROM @drop_query")
	testDB.Exec("EXECUTE stmt")
	testDB.Exec("DEALLOCATE PREPARE stmt")
	testDB.Exec("SET FOREIGN_KEY_CHECKS = 1")
	return nil
}

// 모델 등록 부분
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Board{},
		&model.Document{},
	)
}
