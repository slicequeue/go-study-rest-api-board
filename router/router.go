package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	// 미들 웨어 설정
	e.Pre(middleware.RemoveTrailingSlash())					// trailing slash 제거
	e.Use(middleware.Logger()) 								// 로거 설정
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ 	// CORS 설정
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Validator = NewValidator()	// 검증기 등록
	return e
}