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
		SELECT id, name, author, publication
		FROM books
		WHERE id = $1 
		`,
		id,
	).Scan(
		&book.ID,
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
			&book.ID,
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
        RETURNING id, name, author, publication
        `,
		newBook.Name,
		newBook.Author,
		newBook.Publication,
	).Scan(
		&book.ID,
		&book.Name,
		&book.Author,
		&book.Publication,
	)

	if err != nil {
		return nil, err
	}
	return &book, nil

}

func (r *BookRepository) UpdateBook(ctx context.Context, id int64, book models.Book) (*models.Book, error) {
	var updatedBook models.Book
	err := r.db.QueryRow(
		ctx,
		`
		UPDATE books
		SET name = $1, author = $2, publication = $3
		WHERE id = $4
		RETURNING name, author, publication
		`,
		book.Name,
		book.Author,
		book.Publication,
		id,
	).Scan(
		&updatedBook.Name,
		&updatedBook.Author,
		&updatedBook.Publication,
	)

	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) DeleteBook(ctx context.Context, id int64) error {
	_, err := r.db.Exec(
		ctx,
		`
		DELETE FROM books
		WHERE id = $1
		`,
		id,
	)
	return err
}
