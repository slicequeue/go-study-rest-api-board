package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/slicequeue/go-study-rest-api-board/store"
	"github.com/slicequeue/go-study-rest-api-board/utils"
)

// signin 처리하기
func (h *Handler) SignIn(c echo.Context) error {
	req := &SignInRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	signInDto := store.NewSignInDto(req.Email, req.Password)
	u, t, err := h.authStore.SignIn(signInDto)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newSignInResponse(u, t))
}
