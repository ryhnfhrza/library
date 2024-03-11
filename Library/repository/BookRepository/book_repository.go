package repository

import (
	"context"
	"database/sql"
	domain "library/Model/Domain"
)

type BookRepository interface {
	AddBook(ctx context.Context, tx *sql.Tx,book domain.Book)domain.Book
	UpdateBook(ctx context.Context, tx *sql.Tx,book domain.Book)domain.Book
	DeleteBook(ctx context.Context, tx *sql.Tx,book domain.Book)
	FindBookById(ctx context.Context, tx *sql.Tx,bookId string)(domain.Book,error)
	FindAllBook(ctx context.Context, tx *sql.Tx)[]domain.Book
}