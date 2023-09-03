package store

import (
	"github.com/slicequeue/go-study-rest-api-board/model"
	"gorm.io/gorm"
)

type DocumentStore struct {
	db *gorm.DB
}

func NewDocumentStore(db *gorm.DB) *DocumentStore {
	return &DocumentStore{
		db: db,
	}
}

func (ds *DocumentStore) Create(d *model.Document) (err error) {
	return ds.db.Create(d).Preload("User").Error
}
