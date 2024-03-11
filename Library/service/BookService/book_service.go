package service

import (
	"context"
	web "library/Model/Web/BookWeb"
)

type BookService interface {
	AddBook(ctx context.Context,request web.BookAddRequest)web.BookResponse
	UpdateBook(ctx context.Context,request web.BookUpdateRequest)web.BookResponse
	DeleteBook(ctx context.Context,bookId string)
	FindBookById(ctx context.Context,bookId string)web.BookResponse
	FindAllBook(ctx context.Context)[]web.BookResponse
}