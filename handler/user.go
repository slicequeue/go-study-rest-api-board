package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/slicequeue/go-study-rest-api-board/model"
	"github.com/slicequeue/go-study-rest-api-board/utils"
)

func (h *Handler) SignUp(c echo.Context) error {
	var u model.User
	req := &userRegisterRequest{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	fmt.Println("model.User", u)
	if err := h.userStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newUserResponse(&u))
}
