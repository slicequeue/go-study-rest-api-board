package main

import (
	"fmt"

	"github.com/slicequeue/go-study-rest-api-board/router"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	r := router.New()

	// r.GET("/swagger/*", echoSwagger.WrapHandler) // TODO

	v1 := r.Group("/api") // TODO

	// d := db.New() // TODO
	// db.AutoMigrate(d)
	
	
}