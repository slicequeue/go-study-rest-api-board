package main

import (
	"github.com/slicequeue/go-study-rest-api-board/db"
	"github.com/slicequeue/go-study-rest-api-board/handler"
	"github.com/slicequeue/go-study-rest-api-board/router"
)

func main() {
	r := router.New()

	// r.GET("/swagger/*", echoSwagger.WrapHandler) // TODO
	d := db.New()
	db.AutoMigrate(d)

	v1 := r.Group("/api") // TODO
	h := handler.NewHandler()
	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8585"))

}
