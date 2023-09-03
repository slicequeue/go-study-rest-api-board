package main

import (
	"github.com/slicequeue/go-study-rest-api-board/db"
	"github.com/slicequeue/go-study-rest-api-board/handler"
	"github.com/slicequeue/go-study-rest-api-board/router"
	"github.com/slicequeue/go-study-rest-api-board/store"
)

func main() {
	r := router.New()

	// r.GET("/swagger/*", echoSwagger.WrapHandler) // TODO
	d := db.New()
	db.AutoMigrate(d)

	v1 := r.Group("/api") // TODO
	us := store.NewUserStore(d)
	as := store.NewAuthStore(d, us)
	ds := store.NewDocumentStore(d)
	bs := store.NewBoardStore(d)
	h := handler.NewHandler(as, us, bs, ds)
	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8585"))

}
