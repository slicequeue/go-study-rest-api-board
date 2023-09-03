package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/slicequeue/go-study-rest-api-board/router/middleware"
	"github.com/slicequeue/go-study-rest-api-board/utils"
)

type HealthCheckInfo struct {
	ServerName string    `json:"name"`
	At         time.Time `json:"at"`
}

func healthCheck(c echo.Context) error {
	healthCheckInfo := &HealthCheckInfo{
		ServerName: "go-study-rest-api-board",
		At:         time.Now(),
	}
	return c.JSON(http.StatusOK, healthCheckInfo)
}

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	// 기본 헬스체크 경로 등록
	v1.GET("", healthCheck)
	v1.GET("/health", healthCheck)

	guestUsers := v1.Group("/users")
	guestUsers.POST("", h.SignUp)

	auth := v1.Group("/auth")
	auth.POST("/signin", h.SignIn)

	board := v1.Group("/boards", jwtMiddleware)
	board.GET("", h.GetBoards)
	board.GET("/:boardId", h.GetBoardDetail)
	board.POST("/:boardId/documents", h.CreateBoardDocument)
	board.GET("/:boardId/documents", h.GetBoardDocuments)
}
