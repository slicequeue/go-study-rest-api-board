package store

import (
	"fmt"

	"github.com/slicequeue/go-study-rest-api-board/model"
	"gorm.io/gorm"
)

type BoardStore struct {
	db *gorm.DB
}

func NewBoardStore(db *gorm.DB) *BoardStore {
	return &BoardStore{
		db: db,
	}
}

func (bs *BoardStore) GetAll() ([]model.Board, error) {
	var boards []model.Board
	// TODO .Preload("Documents") 부분 추후 SQL 자체 실행으로 성능 강화
	if err := bs.db.Model(&model.Board{}).Preload("User").Preload("Documents").Find(&boards).Error; err != nil {
		return nil, err
	}
	return boards, nil
}

func (bs *BoardStore) GetById(id int) (*model.Board, error) {
	var m model.Board
	if err := bs.db.Preload("User").Preload("Documents").First(&m, id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (bs *BoardStore) Create(b *model.Board) error {
	return bs.db.Model(b).Create(b).Error
}

func (bs *BoardStore) Update(b *model.Board) error {
	return bs.db.Model(b).Updates(b).Error
}

func (bs *BoardStore) AddDocument(b *model.Board, d *model.Document) error {
	(*b).AddDocument(d)
	return bs.Update(b)
}

func (bs *BoardStore) GetBoardDocument(boardId uint, documentId uint) (*model.Document, error) {
	var d model.Document
	if err := bs.db.Preload("User").Where("board_id = ? AND id = ?", boardId, documentId).First(&d).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}

// --- board documents --- //
type PageInfo struct {
	Page          uint
	Size          uint
	IsFirst       bool
	IsLast        bool
	TotalPages    uint
	TotalElements uint
}

func (bs *BoardStore) GetBoardDocuments(boardId int, page int, size int) ([]*model.Document, *PageInfo, error) {
	d := model.Document{}
	var count int64
	if err := bs.db.Model(&d).Where("board_id").Count(&count).Error; err != nil {
		return nil, nil, err
	}

	offset := (page - 1) * size
	limit := size
	var documents []*model.Document
	if err := bs.db.Preload("User").Limit(limit).Offset(offset).Where("board_id = ?", boardId).Find(&documents).Error; err != nil {
		return nil, nil, err
	}
	fmt.Println(documents)

	totalPages := uint(count / int64(size))

	pageDto := PageInfo{
		Page:          uint(page),
		Size:          uint(size),
		IsFirst:       page == 1,
		IsLast:        totalPages == uint(page),
		TotalPages:    totalPages,
		TotalElements: uint(count),
	}

	return documents, &pageDto, nil
}
