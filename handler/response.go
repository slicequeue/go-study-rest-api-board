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

type BoardDto struct {
	Id            uint    `json:"id"`
	Name          string  `json:"name"`
	Author        UserDto `json:"author"`
	DocumentCount int     `json:"documentCount"`
}

type BoardListResponse struct {
	Boards []BoardDto `json:"boards"`
}

func newBoardListResponse(boards []model.Board) *BoardListResponse {
	var boardDtos []BoardDto
	for _, b := range boards {
		// author := UserDto {
		// 	Email: b.User.Email,
		// 	Username: b.User.Username,
		// }
		// boardDto := BoardDto{
		// 	Id: b.ID,
		// 	Name: b.Name,
		// 	Author: author,
		// }
		boardDto := new(BoardDto)
		boardDto.Id = b.ID
		boardDto.Name = b.Name
		boardDto.Author.Email = b.User.Email
		boardDto.Author.Username = b.User.Username
		boardDto.DocumentCount = len(b.Documents)
		boardDtos = append(boardDtos, *boardDto)
	}
	boardListResponse := new(BoardListResponse)
	boardListResponse.Boards = boardDtos
	return boardListResponse
}
