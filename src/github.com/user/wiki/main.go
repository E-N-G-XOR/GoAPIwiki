package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/user/wiki/driver"
	"github.com/user/wiki/handler/http"
)

func main() {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	println("this is db", dbName, dbHost, dbPass, dbPort)

	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewPostHandler(connection)
	r.Get("/posts", pHandler.Fetch)
	r.Get("/posts/{id}", pHandler.GetByID)
	r.Post("/posts/create", pHandler.Create)
	r.Put("/posts/update/{id}", pHandler.Update)
	r.Delete("/posts/{id}", pHandler.Delete)

	fmt.Println("Server listen at :6600")
	http.ListenAndServe(":6600", r)
}
