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
	assert.EqualValues(t, resultUser.ID, user.ID)
	assert.EqualValues(t, resultUser.Email, user.Email)
	assert.EqualValues(t, resultUser.Password, user.Password)

	// Update
	changePassword := "4321"
	changeEmail := "slicequeue@gmail.com"
	user.Password = changePassword
	user.Email = changeEmail
	updateResult := d.Updates(&user)
	assert.NoError(t, updateResult.Error)
	assert.NotEqualValues(t, user.Password, changePassword)
	assert.EqualValues(t, user.Email, changeEmail)

	// Softly Delete
	softDeleteResult1 := d.Delete(&user) // 모델을 포함하고 있기에 부드러운 삭제 가능
	assert.NoError(t, softDeleteResult1.Error)
	softDeleteResult2 := d.Where("id = ?", user.ID).First(&user)
	assert.EqualError(t, softDeleteResult2.Error, gorm.ErrRecordNotFound.Error())
	softDeleteResult3 := d.Unscoped().Where("id = ?", user.ID).First(&user)
	assert.NoError(t, softDeleteResult3.Error)
	assert.EqualValues(t, user.Email, changeEmail)

	// Hard Delete
	hardDeleteResult1 := d.Unscoped().Delete(&user) // Unscoped() 이용하여 삭제 가능
	assert.NoError(t, hardDeleteResult1.Error)
	hardDeleteResult2 := d.Where("id = ?", user.ID).First(&user)
	assert.EqualError(t, hardDeleteResult2.Error, gorm.ErrRecordNotFound.Error())

	tearDown()

}
