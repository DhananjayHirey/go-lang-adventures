package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/repositories"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BookController struct {
	bookRepository *repositories.BookRepository
}

func NewBookController(
	bookRepository *repositories.BookRepository,
) *BookController {
	return &BookController{
		bookRepository: bookRepository,
	}
}

func (c *BookController) GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.Atoi(params["bookId"])
	bookId64 := int64(bookId)
	if err != nil {
		fmt.Errorf("Failed to convert bookId %w", err)
		return
	}
	book, err := c.bookRepository.GetBookById(r.Context(), bookId64)

	if err != nil {
		http.Error(
			w,
			"failed to get book",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(book)

}

func (c *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	books, err := c.bookRepository.GetBook(r.Context())
	if err != nil {
		http.Error(w, "failed to get books", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	_ = json.NewDecoder(r.Body).Decode(&newBook)
	createdBook, err := c.bookRepository.CreateBook(r.Context(), newBook)
	if err != nil {
		http.Error(w, "failed to create books", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdBook)
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var bookUpdates models.Book
	_ = json.NewDecoder(r.Body).Decode(&bookUpdates)
	updatedBook, err := c.bookRepository.UpdateBook(r.Context(), bookUpdates.ID, bookUpdates)
	if err != nil {
		http.Error(w, "failed to update Books", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedBook)
}

func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.Atoi(params["bookId"])
	bookId64 := int64(bookId)
	if err != nil {
		fmt.Errorf("Failed to convert bookId %w", err)
		return
	}

	err2 := c.bookRepository.DeleteBook(r.Context(), bookId64)

	if err2 != nil {
		http.Error(w, "failed to delete book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
