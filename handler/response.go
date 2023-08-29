package handler

import (
	"github.com/slicequeue/go-study-rest-api-board/model"
)

type UserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserResponse struct {
	User UserDto `json:"user"`
}

func newUserResponse(u *model.User) *UserResponse {
	r := new(UserResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	return r
}

type SignInResponse struct {
	User  UserDto `json:"user"`
	Token string  `json:"token"`
}

func newSignInResponse(u *model.User, token string) *SignInResponse {
	r := new(SignInResponse)
	r.User.Username = (*u).Username
	r.User.Email = (*u).Email
	r.Token = token
	return r
}
