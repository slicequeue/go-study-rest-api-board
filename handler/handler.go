package handler

import "github.com/slicequeue/go-study-rest-api-board/store"

// 헨들러 struct 등록하기
type Handler struct {
	userStore *store.UserStore
}

// 헨들러 받아서 설정하기
func NewHandler(userStore *store.UserStore) *Handler {
	return &Handler{userStore: userStore}
}
