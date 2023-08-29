package store

import (
	"errors"

	"github.com/slicequeue/go-study-rest-api-board/model"
	"github.com/slicequeue/go-study-rest-api-board/utils"
	"gorm.io/gorm"
)

type AuthStore struct {
	db *gorm.DB
	us *UserStore
}

func NewAuthStore(db *gorm.DB, us *UserStore) *AuthStore {
	return &AuthStore{
		db: db,
		us: us,
	}
}

type SignInDto struct {
	Email    string
	Password string
}

func NewSignInDto(email, password string) *SignInDto {
	return &SignInDto{
		Email:    email,
		Password: password,
	}
}

func (as *AuthStore) SignIn(signInDto *SignInDto) (*model.User, string, error) {
	user, err := as.us.GetByEmail(signInDto.Email)
	if err != nil {
		return nil, "", err
	}
	if user == nil {
		return nil, "", errors.New("user not exist!")
	}
	if !user.CheckPassword(signInDto.Password) {
		return nil, "", errors.New("wrong password!")
	}
	token := utils.GenerateJWT(user.ID)
	return user, token, nil
}
