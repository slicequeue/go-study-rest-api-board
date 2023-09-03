package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/slicequeue/go-study-rest-api-board/model"
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

func (h *Handler) GetBoardDetail(c echo.Context) error {
	boardId := c.Param("boardId") // getParam(c, "boardId", "integer")
	id, err := strconv.Atoi(boardId)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid boardId")
	}
	board, err := h.boardStore.GetById(id)
	if err != nil {

	}
	return c.JSON(http.StatusOK, newBoardDetailResponse(board))
}

var (
	DEFAULT_BOARD_DOCUMENTS_SIZE int = 10
	DEFAULT_BOARD_DOCUMENTS_PAGE int = 1
)

func (h *Handler) GetBoardDocuments(c echo.Context) error {
	boardId := getParamValue(c, "boardId", "integer", nil).(int)
	page := getQueryParam(c, "page", "integer", DEFAULT_BOARD_DOCUMENTS_PAGE).(int)
	size := getQueryParam(c, "size", "integer", DEFAULT_BOARD_DOCUMENTS_SIZE).(int)

	board, _ := h.boardStore.GetById(boardId)
	if board == nil {
		return c.String(http.StatusNotFound, "board not found.")
	}

	boardDocuments, pageInfo, err := h.boardStore.GetBoardDocuments(boardId, page, size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newBoardDocumentsResponse(board, boardDocuments, pageInfo))
}

func (h *Handler) CreateBoardDocument(c echo.Context) error {
	var document model.Document
	req := new(BoardDocumentRequest)
	if err := req.bind(c, &document); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	fmt.Println("model.Document", document)
	board, boardErr := h.boardStore.GetById(int(document.BoardID))
	if boardErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(boardErr))
	}
	if board == nil {
		return c.String(http.StatusNotFound, "board not found")
	}
	if err := h.documentStore.Create(&document); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	// 아래 코드 이슈가 있음! - DB에 삽입까지는 잘 되나 &document 넘겨준 객체에 PK 값이 제대로 나오지 않음..
	// if err := h.boardStore.AddDocument(board, &document); err != nil {
	// 	return c.JSON(http.StatusUnprocessableEntity, utils.NewError(boardErr))
	// }
	// TODO 너저분하다 Create 이후 User Preload 가 되지 않아 직접 조회하는 형태라니...
	resultDocument, documentFoundErr := h.boardStore.GetBoardDocument(board.ID, document.ID)
	if documentFoundErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(documentFoundErr))
	}
	fmt.Println("resultDocument", resultDocument)
	return c.JSON(http.StatusOK, newDocumentResponse(resultDocument))
}

// --- private --- //

func getParamValue(c echo.Context, key string, typeName string, defaultValue interface{}) interface{} {
	value := c.Param(key)
	convertedValue, err := convertTypeValue(typeName, value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	if convertedValue == nil {
		return defaultValue
	}
	return convertedValue
}

func getQueryParam(c echo.Context, key string, typeName string, defaultValue interface{}) interface{} {
	value := c.QueryParam(key)
	convertedValue, err := convertTypeValue(typeName, value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	if convertedValue == nil {
		return defaultValue
	}
	return convertedValue
}

func convertTypeValue(typeName string, value string) (interface{}, error) {
	if value == "" {
		return nil, nil
	}
	switch typeName {
	case "string":
		return value, nil
	case "integer":
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return intValue, nil
	default:
		return nil, errors.New("Unsupported type")
	}
}
