package app

import (
	"HumosBooks/DataBase"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)
type UserSvc struct {
	DB *sql.DB
}
func NewUserSvc(DB *sql.DB) *UserSvc {
	return &UserSvc{DB: DB}
}

type MainServer struct {
	DB *sql.DB
	Router *httprouter.Router
	UserSvc *UserSvc
}

func NewMainServer(DB *sql.DB, router *httprouter.Router, userSvc *UserSvc) *MainServer {
	return &MainServer{DB: DB, Router: router, UserSvc: userSvc}
}

func (server *MainServer) Start() {
	DataBase.DbInit(server.DB)
	server.InitRoutes()
}

func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.Router.ServeHTTP(writer, request)
}

/*
package services

import (
	"database/sql"
)

type UserSvc struct {
	DB *sql.DB
}
func NewUserSvc(DB *sql.DB) *UserSvc {
	return &UserSvc{DB: DB}
}
*/