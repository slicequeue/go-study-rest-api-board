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