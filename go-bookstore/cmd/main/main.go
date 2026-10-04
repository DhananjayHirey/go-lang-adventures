package main

import (
	"context"
	"go-bookstore/pkg/config"
	"go-bookstore/pkg/controllers"
	"go-bookstore/pkg/repositories"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	ctx := context.Background()

	// db pool connection
	db, err := config.NewPostgresPool(
		ctx,
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bookRepo := repositories.NewBookRepository(db)
	bookController := controllers.NewBookController(bookRepo)

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r, bookController)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
