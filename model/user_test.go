package model_test

import (
	"fmt"
	"testing"

	"github.com/slicequeue/go-study-rest-api-board/db"
	"github.com/slicequeue/go-study-rest-api-board/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	d *gorm.DB
)

func setUp() {
	d = db.TestDB()
	db.AutoMigrate(d)
}

func tearDown() {
	db.DropTestDB()
	d, _ := d.DB()
	d.Close()
}

func TestUserCRUD(t *testing.T) {
	setUp()
	// Create
	user := model.User{
		Email:    "gskjhyoo@naver.com",
		Username: "slicequeue",
		Password: "1234",
	}
	fmt.Println("prev insert user:", user)
	createResult := d.Create(&user)
	assert.NoError(t, createResult.Error)
	fmt.Println("after insert user:", user)

	// Read
	var resultUser model.User;
	readResult := d.Where("id = ?", user.ID).First(&resultUser)
	assert.NoError(t, readResult.Error)

	// TODO: Update

	// TODO: Softly Delete

	// TODO: Hard Delete

	tearDown()

}
