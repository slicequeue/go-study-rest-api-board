package store

import (
	"github.com/slicequeue/go-study-rest-api-board/model"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetById(id uint) (*model.User, error) {
	var m model.User
	if err := us.db.First(&m, id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) GetByEmail(email string) (*model.User, error) {
	var m model.User
	if err := us.db.Where("email = ?", email).First(&m).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) Create(u *model.User) (err error) {
	return us.db.Create(u).Error
}

func (us *UserStore) Update(u *model.User) error {
	return us.db.Model(u).Updates(u).Error
}

func (us *UserStore) AddFollower(u *model.User, followerID uint) error {
	return us.db.Model(u).Association("Followers").Append(&model.Follow{FollowerID: followerID, FollowingID: u.ID})
}
