package service

import (
	"context"
	"database/sql"
	domain "library/Model/Domain"
	web "library/Model/Web/BookWeb"
	"library/exception"
	"library/helper"
	repository "library/repository/BookRepository"

	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct{
	BookRepository repository.BookRepository
	Db *sql.DB
	Validate *validator.Validate
}

func NewBookService(bookRepository repository.BookRepository,db *sql.DB,validate *validator.Validate)BookService{
	return &BookServiceImpl{
		BookRepository: bookRepository,
		Db: db,
		Validate: validate,
	}
}

func(service *BookServiceImpl)AddBook(ctx context.Context,request web.BookAddRequest)web.BookResponse{
	err:=service.Validate.Struct(request)
	helper.PanicIfError(err)

 tx,err := service.Db.Begin()
 helper.PanicIfError(err)
 defer helper.CommitOrRollback(tx)

 book := domain.Book{
	Id: request.Id,
	Title: request.Title,
	Author: request.Author,
	Category: request.Category,
 }

 book = service.BookRepository.AddBook(ctx,tx,book)

 return helper.ToBookResponse(book)
}

func(service *BookServiceImpl)UpdateBook(ctx context.Context,request web.BookUpdateRequest)web.BookResponse{
	err:=service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book,err := service.BookRepository.FindBookById(ctx,tx,request.Id)	
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	request.Title = helper.GetDefaultIfEmpty(request.Title,book.Title)
	request.Author = helper.GetDefaultIfEmpty(request.Author,book.Author)
	request.Category = helper.GetDefaultIfEmpty(request.Category,book.Category)
	
	book.Title = request.Title
	book.Author = request.Author
	book.Category = request.Category

	 book = service.BookRepository.UpdateBook(ctx,tx,book)

	 return helper.ToBookResponse(book)
}

func(service *BookServiceImpl)DeleteBook(ctx context.Context,bookId string){
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	 defer helper.CommitOrRollback(tx)
 
	book,err := service.BookRepository.FindBookById(ctx,tx,bookId)	
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BookRepository.DeleteBook(ctx,tx,book)
}

func(service *BookServiceImpl)FindBookById(ctx context.Context,bookId string)web.BookResponse{
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
 
	book,err := service.BookRepository.FindBookById(ctx,tx,bookId)	
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBookResponse(book)
}

func(service *BookServiceImpl)FindAllBook(ctx context.Context)[]web.BookResponse{
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	books:= service.BookRepository.FindAllBook(ctx,tx)

	return helper.ToBookResponses(books)
}
