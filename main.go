package main

import (
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	//"HumosBooks/pkg/core/services"
	"HumosBooks/cmd/app"
)
func main() {
	database, err := sql.Open("sqlite3", "db")
	if err != nil {
		fmt.Println("Connection with DB. Error is", err)
	}
	router := httprouter.New()
	svc := /*services.*/app.NewUserSvc(database)
	server := app.NewMainServer(database, router, svc)
	fmt.Println(" server starting ")
	server.Start()
	fmt.Println(" server started ")

}