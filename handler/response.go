package handler

import (
	"time"

	"github.com/slicequeue/go-study-rest-api-board/model"
	"github.com/slicequeue/go-study-rest-api-board/store"
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
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Author    UserDto   `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type BoardListResponse struct {
	Boards []BoardDto `json:"boards"`
}

func newBoardListResponse(boards []model.Board) *BoardListResponse {
	var boardDtos []BoardDto
	for _, b := range boards {
		// 이 방식도 되고
		// author := UserDto {
		// 	Email: b.User.Email,
		// 	Username: b.User.Username,
		// }
		// boardDto := BoardDto{
		// 	Id: b.ID,
		// 	Name: b.Name,
		// 	Author: author,
		// }
		// 이 방식도 됨!
		boardDto := new(BoardDto)
		boardDto.Id = b.ID
		boardDto.Name = b.Name
		boardDto.Author.Email = b.User.Email
		boardDto.Author.Username = b.User.Username
		boardDtos = append(boardDtos, *boardDto)
		boardDto.CreatedAt = b.CreatedAt
		boardDto.UpdatedAt = b.UpdatedAt
	}
	boardListResponse := new(BoardListResponse)
	boardListResponse.Boards = boardDtos
	return boardListResponse
}

type BoardDetailResponse struct {
	Board BoardDto `json:"board"`
	Info  struct {
		DocumentCount int `json:"documentCount"`
	} `json:"info"`
}

func newBoardDetailResponse(board *model.Board) *BoardDetailResponse {
	boardDetailResponse := new(BoardDetailResponse)
	// boardDetailResponse.Board = &BoardDto{}
	boardDetailResponse.Board.Id = board.ID
	boardDetailResponse.Board.Name = board.Name
	boardDetailResponse.Board.Author.Username = board.User.Username
	boardDetailResponse.Board.Author.Email = board.User.Email
	boardDetailResponse.Board.CreatedAt = board.CreatedAt
	boardDetailResponse.Board.UpdatedAt = board.UpdatedAt
	boardDetailResponse.Info.DocumentCount = len(board.Documents)
	return boardDetailResponse
}

type BoardDocumentsResponseDto struct {
	Board     BoardDto      `json:"board"`
	Documents []DocumentDto `json:"documents"`
	Page      PageDto       `json:"page"`
}

type DocumentDto struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    UserDto   `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PageDto struct {
	Page          uint `json:"page"`
	Size          uint `json:"size"`
	IsFirst       bool `json:"isFirst"`
	IsLast        bool `json:"isLast"`
	TotalPages    uint `json:"totalPages"`
	TotalElements uint `json:"totalElements"`
}

func newBoardDocumentsResponse(board *model.Board, boardDocuments []*model.Document, pageInfo *store.PageInfo) *BoardDocumentsResponseDto {
	boardDocumentsResponseDto := new(BoardDocumentsResponseDto)
	boardDocumentsResponseDto.Board.Id = board.ID
	boardDocumentsResponseDto.Board.Name = board.Name
	boardDocumentsResponseDto.Board.Author.Username = board.User.Username
	boardDocumentsResponseDto.Board.Author.Email = board.User.Email
	boardDocumentsResponseDto.Board.CreatedAt = board.CreatedAt
	boardDocumentsResponseDto.Board.UpdatedAt = board.UpdatedAt
	boardDocumentsResponseDto.Page.Page = pageInfo.Page
	boardDocumentsResponseDto.Page.Size = pageInfo.Size
	boardDocumentsResponseDto.Page.IsFirst = pageInfo.IsFirst
	boardDocumentsResponseDto.Page.IsLast = pageInfo.IsLast
	boardDocumentsResponseDto.Page.TotalPages = pageInfo.TotalPages
	boardDocumentsResponseDto.Page.TotalElements = pageInfo.TotalElements

	docuements := []DocumentDto{}
	for _, boardDocument := range boardDocuments {
		document := new(DocumentDto)
		document.Id = boardDocument.ID
		document.Title = boardDocument.Title
		document.Author.Username = boardDocument.Author.User.Username
		document.Author.Email = boardDocument.Author.User.Email
		document.CreatedAt = boardDocument.RecordAt.CreatedAt
		document.UpdatedAt = boardDocument.RecordAt.UpdatedAt
		docuements = append(docuements, *document)
	}
	boardDocumentsResponseDto.Documents = docuements

	return boardDocumentsResponseDto
}

func newDocumentResponse(document *model.Document) *DocumentDto {
	documentDto := new(DocumentDto)
	documentDto.Id = document.ID
	documentDto.Title = document.Title
	documentDto.Content = document.Content
	documentDto.Author.Email = document.Author.User.Email
	documentDto.Author.Username = document.Author.User.Username
	documentDto.CreatedAt = document.RecordAt.CreatedAt
	documentDto.UpdatedAt = document.RecordAt.UpdatedAt
	return documentDto
}
