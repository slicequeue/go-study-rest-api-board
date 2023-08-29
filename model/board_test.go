package model_test

import (
	"testing"

	"github.com/slicequeue/go-study-rest-api-board/db"
	"github.com/slicequeue/go-study-rest-api-board/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBoardCRUD(t *testing.T) {
	setUp()
	user := model.User{
		Email:    "slicequeue@gmail.com",
		Username: "slicequeue",
		Password: "1234",
	}
	d.Create(&user)

	// Create 1 - UserId, User 둘다 세팅
	board1 := model.Board{
		Name:   "보드1",
		UserID: user.ID,
		User:   user,
	}
	createResult1 := d.Create(&board1)
	assert.NotNil(t, board1.ID)
	assert.NoError(t, createResult1.Error)

	// Create 2 - UserId 만 세팅
	board2 := model.Board{
		Name:   "보드2",
		UserID: user.ID,
	}
	createResult2 := d.Create(&board2)
	assert.NoError(t, createResult2.Error)

	// Create 3 - User 만 세팅
	board3 := model.Board{
		Name: "보드3",
		User: user,
	}
	createResult3 := d.Create(&board3)
	assert.NoError(t, createResult3.Error)

	board4 := model.Board{
		Name: "보드4",
	}
	createResult4 := d.Create(&board4)
	assert.Error(t, createResult4.Error)

	// Read 1 - ID 조회
	var foundBoard1 model.Board
	readResult1 := d.Where("id = ?", board1.ID).First(&foundBoard1)
	assert.EqualValues(t, foundBoard1.ID, board1.ID)
	assert.EqualValues(t, foundBoard1.Name, board1.Name)
	assert.EqualValues(t, foundBoard1.UserID, board1.UserID)
	assert.NotEqualValues(t, foundBoard1.User.ID, board1.User.ID) // 이 부분 Lazy Loading 이 되지 않으므로 기본값으로 같지 않음
	assert.EqualValues(t, foundBoard1.User.ID, uint(0))           // 기본 값으로 같다
	assert.NoError(t, readResult1.Error)

	// Read 2 - ID 조회 & Eager Loading
	var foundBoard2 model.Board
	readResult2 := d.Where("id = ?", board2.ID).Preload("User").First(&foundBoard2)
	assert.NoError(t, readResult2.Error)
	assert.EqualValues(t, foundBoard2.ID, board2.ID)
	assert.EqualValues(t, foundBoard2.Name, board2.Name)
	assert.EqualValues(t, foundBoard2.UserID, board2.UserID)
	assert.NotEqualValues(t, foundBoard2.User.ID, board2.User.ID)

	// Update
	changeName := "이름변경한보드2"
	board2.Name = changeName
	updateResult2 := d.Updates(&board2)
	assert.NoError(t, updateResult2.Error)
	assert.EqualValues(t, board2.Name, changeName)

	// Softly Delete
	var foundBoard3 model.Board
	softDeleteResult3 := d.Delete(&board3)
	assert.NoError(t, softDeleteResult3.Error)
	softDeleteFoundResult3 := d.Where("id = ?", board3.ID).First(&foundBoard3)
	assert.EqualError(t, softDeleteFoundResult3.Error, gorm.ErrRecordNotFound.Error())

	// Hard Delete
	var foundBoard4 model.Board
	hardDeleteResult4 := d.Unscoped().Where("id = ?", board4.ID).Delete(&board4)
	assert.NoError(t, hardDeleteResult4.Error)
	hardDeleteFoundResult4 := d.Where("id = ?", board4.ID).First(&foundBoard4)
	assert.EqualError(t, hardDeleteFoundResult4.Error, gorm.ErrRecordNotFound.Error())

	tearDown()
}

func TestDocumentCRUD(t *testing.T) {
	setUp()
	user := model.User{
		Email:    "slicequeue@gmail.com",
		Username: "slicequeue",
		Password: "1234",
	}
	d.Create(&user)
	board1 := model.Board{
		Name:   "보드1",
		UserID: user.ID,
		User:   user,
	}
	d.Create(&board1)

	author := model.Author{
		User: user,
	}

	// Create
	document1 := model.Document{
		Board:   board1,
		Title:   "글1",
		Content: "글1내용",
		Author:  author,
	}
	createResult1 := d.Create(&document1)
	assert.NoError(t, createResult1.Error)

	// Read
	var foundDocument1 model.Document
	readResult1 := d.Preload("Author.User").Where("id = ?", document1.ID).First(&foundDocument1)
	assert.NoError(t, readResult1.Error)
	assert.NotNil(t, foundDocument1.ID)
	assert.EqualValues(t, foundDocument1.Title, document1.Title)
	assert.EqualValues(t, foundDocument1.BoardID, document1.BoardID)
	assert.EqualValues(t, foundDocument1.Content, document1.Content)
	assert.EqualValues(t, foundDocument1.Author.UserId, document1.Author.UserId)
	assert.EqualValues(t, foundDocument1.Author.User.ID, document1.Author.User.ID)
	assert.EqualValues(t, foundDocument1.RecordAt.CreatedAt, document1.RecordAt.CreatedAt)
	assert.EqualValues(t, foundDocument1.RecordAt.UpdatedAt, document1.RecordAt.UpdatedAt)
	assert.EqualValues(t, foundDocument1.RecordAt.DeletedAt, document1.RecordAt.DeletedAt)

	// Update
	changeContent := "글1내용변경"
	document1.Content = changeContent
	updateResult1 := d.Updates(&document1)
	assert.NoError(t, updateResult1.Error)
	assert.EqualValues(t, document1.Content, changeContent)

	// soft Delete
	deleteResult1 := d.Delete(&document1)
	assert.NoError(t, deleteResult1.Error)
	var deleteFoundDocument model.Document
	deleteFoundResult := d.Where("id = ?", document1.ID).First(&deleteFoundDocument)
	assert.EqualError(t, deleteFoundResult.Error, gorm.ErrRecordNotFound.Error())

	tearDown()
}

func TestBoardDocumentAssociation_case1(t *testing.T) {
	setUp()
	user := model.User{
		Email:    "slicequeue@gmail.com",
		Username: "slicequeue",
		Password: "1234",
	}
	d.Create(&user)
	board1 := model.Board{
		Name:   "보드1",
		UserID: user.ID,
		User:   user,
	}
	d.Create(&board1)

	author := model.Author{
		User: user,
	}

	document1 := model.Document{
		Board:   board1,
		Title:   "글1",
		Content: "글1내용",
		Author:  author,
	}
	document2 := model.Document{
		Board:   board1,
		Title:   "글2",
		Content: "글2내용",
		Author:  author,
	}
	document3 := model.Document{
		Board:   board1,
		Title:   "글2",
		Content: "글2내용",
		Author:  author,
	}

	board1.Documents = []model.Document{
		document1,
		document2,
		document3,
	}

	// Create with association
	d.Updates(&board1)
	tearDown() 
}
