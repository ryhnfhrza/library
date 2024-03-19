package repository

import (
	"context"
	"database/sql"
	"errors"
	domain "library/Model/Domain"
	"library/helper"
)

type BookRepositoryImpl struct{

}

func NewBookRepository()BookRepository{
	return &BookRepositoryImpl{}
}

func(repository *BookRepositoryImpl)AddBook(ctx context.Context, tx *sql.Tx,book domain.Book)domain.Book{
	SQL := "insert into book (id,title,author,category) values (?,?,?,?)"
	_,err:=tx.ExecContext(ctx,SQL,book.Id,book.Title,book.Author,book.Category)
	helper.PanicIfError(err)

	return book
}

func(repository *BookRepositoryImpl)UpdateBook(ctx context.Context, tx *sql.Tx,book domain.Book)domain.Book{
	SQL := "update book set title = ? , author = ?, category = ? where id = ?"
	_,err:=tx.ExecContext(ctx,SQL,book.Title,book.Author,book.Category, book.Id)
	helper.PanicIfError(err)

	return book
}

func(repository *BookRepositoryImpl)DeleteBook(ctx context.Context, tx *sql.Tx,book domain.Book){
	SQL := "delete from book where id = ?"
	_,err:=tx.ExecContext(ctx,SQL,book.Id)
	helper.PanicIfError(err)

}

func(repository *BookRepositoryImpl)FindBookById(ctx context.Context, tx *sql.Tx,bookId string)(domain.Book,error){
	SQL := "select id,title,author,category from book where id = ?"
	rows,err := tx.QueryContext(ctx,SQL,bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	book := domain.Book{}
	if rows.Next(){
		err := rows.Scan(&book.Id,&book.Title,&book.Author,&book.Category)
		helper.PanicIfError(err)
		return book , nil

	}else{
		return book,errors.New("book id "+ bookId+" not found")
	}
}

func(repository *BookRepositoryImpl)FindAllBook(ctx context.Context, tx *sql.Tx)[]domain.Book{
	SQL := "select id,title,author,category from book"
	rows,err := tx.QueryContext(ctx,SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var books []domain.Book
	for rows.Next(){
		book := domain.Book{}
		err := rows.Scan(&book.Id,&book.Title,&book.Author,&book.Category)
		helper.PanicIfError(err)
		books = append(books, book)

	}
	return books
}
