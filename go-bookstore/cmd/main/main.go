package main

import (
	"context"
	"fmt"
	"go-bookstore/pkg/config"
	"go-bookstore/pkg/controllers"
	"go-bookstore/pkg/repositories"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
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
	fmt.Println("Starting the server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
