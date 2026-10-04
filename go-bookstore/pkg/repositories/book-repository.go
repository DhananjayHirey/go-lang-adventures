package repositories

import (
	"context"
	"go-bookstore/pkg/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BookRepository struct {
	db *pgxpool.Pool
}

func NewBookRepository(db *pgxpool.Pool) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) GetBookById(
	ctx context.Context,
	id int64,
) (*models.Book, error) {
	var book models.Book

	err := r.db.QueryRow(
		ctx,
		`
		SELECT name, author, publication
		FROM books
		WHERE id = $1 
		`,
		id,
	).Scan(
		&book.Name,
		&book.Author,
		&book.Publication,
	)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) GetBook(ctx context.Context) (*[]models.Book, error) {
	var books []models.Book

	rows, err := r.db.Query(ctx, `SELECT * FROM books`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.Name,
			&book.Author,
			&book.Publication,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &books, nil

}

func (r *BookRepository) CreateBook(ctx context.Context, newBook models.Book) (*models.Book, error) {
	var book models.Book

	err := r.db.QueryRow(
		ctx,
		`
        INSERT INTO books (name, author, publication)
        VALUES ($1, $2, $3)
        RETURNING name, author, publication
        `,
		newBook.Name,
		newBook.Author,
		newBook.Publication,
	).Scan(
		&book.Name,
		&book.Author,
		&book.Publication,
	)

	if err != nil {
		return nil, err
	}
	return &book, nil

}
