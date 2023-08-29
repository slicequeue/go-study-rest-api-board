package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/slicequeue/go-study-rest-api-board/utils"
)

func (h *Handler) GetBoards(c echo.Context) error {
	boards, err := h.boardStore.GetAll()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	fmt.Println("boards", boards)
	return c.JSON(http.StatusOK, newBoardListResponse(boards))
}
