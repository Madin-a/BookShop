package app

import (
	"HumosBooks/middlware"
	"fmt"
	"log"
	"net/http"
)

func (server *MainServer) InitRoutes() {
	fmt.Println("Routes are init in localhost: 8888")

	server.Router.POST("/api/humo/registration", server.RegistrationHandler)
	server.Router.POST("/api/humo/login", server.LoginHandler)
	server.Router.POST("/api/humo/books/add", middlware.Authorized()(middlware.IsAdmin()(server.AddBookHandler)))
	server.Router.DELETE("/api/humo/books", middlware.Authorized()(middlware.IsAdmin()(server.DeleteBookByHandler)))
	server.Router.GET("/api/humo/books/", middlware.Authorized()(server.SearchBookHandler))
	server.Router.POST("/api/humo/books/buy", middlware.Authorized()(server.BooksBuyHandler))
	server.Router.GET("/api/humo/books/byID", middlware.Authorized()(server.SearchBookByIDHandler))

	err := http.ListenAndServe(":8888", server)
	if err != nil {
		log.Fatal("Can't listen port 8888, Error is", err)
	}

}