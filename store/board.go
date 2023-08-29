package store

import (
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

func (bs *BoardStore) GetById(id uint) (*model.Board, error) {
	var m model.Board
	if err := bs.db.First(&m, id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (bs *BoardStore) Create(b *model.Board) error {
	return bs.db.Model(b).Updates(b).Error
}

func (bs *BoardStore) Update(b *model.Board) error {
	return bs.db.Model(b).Updates(b).Error
}

func (bs *BoardStore) AddDocument(b *model.Board, d *model.Document) error {
	(*b).AddDocument(d)
	return bs.Update(b)
}
