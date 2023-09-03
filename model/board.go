package model

import (
	"time"

	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	Name      string `gorm:"size:256"`
	UserID    uint
	User      User `gorm:"references:ID;not null"`
	Documents []Document
}

func (b *Board) AddDocument(d *Document) {
	if b.Documents == nil {
		b.Documents = []Document{}
	}
	b.Documents = append(b.Documents, *d)
}

type Document struct {
	ID       uint     `gorm:"primarykey;autoIncrement"`
	BoardID  uint     
	Board    Board    `gorm:"references:ID;not null"`
	Title    string   `gorm:"size:256;not null"`
	Content  string   `gorm:"not null"`
	Author   Author   `gorm:"embedded"`
	RecordAt RecordAt `gorm:"embedded"`
}

type Author struct {
	UserId uint
	User   User `gorm:"references:ID;not null"`
}

type RecordAt struct {
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
