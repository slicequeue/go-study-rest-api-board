package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
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
	// 기본 헬스체크 경로 등록
	v1.GET("", healthCheck)
	v1.GET("/health", healthCheck)

	guestUsers := v1.Group("/users")
	guestUsers.POST("", h.SignUp)

	auth := v1.Group("/auth")
	auth.POST("/signin", h.SignIn)

	board := v1.Group("/boards")
	board.GET("", h.GetBoards)
}
