package handler

import "github.com/slicequeue/go-study-rest-api-board/store"

// 헨들러 struct 등록하기
type Handler struct {
	authStore *store.AuthStore
	userStore *store.UserStore
	boardStore *store.BoardStore
}

// 헨들러 받아서 설정하기
func NewHandler(authStore *store.AuthStore, userStore *store.UserStore, boardStore *store.BoardStore) *Handler {
	return &Handler{
		authStore: authStore,
		userStore: userStore,
		boardStore: boardStore,
	}
}
