package main

import (
	"library/app"
	bookController "library/controller/BookController"
	memberController "library/controller/MemberController"
	transactionController "library/controller/TransactionController"
	"library/exception"
	"library/helper"
	bookRepository "library/repository/BookRepository"
	memberRepository "library/repository/MemberRepository"
	transactionRepository "library/repository/TransactionRepository"
	bookService "library/service/BookService"
	memberService "library/service/MemberService"
	transactionService "library/service/TransactionService"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main(){

	//db 
	db := app.NewDb()

	//validate
	validate := validator.New()

	memberRepositroy := memberRepository.NewMemberRepository()
	memberService := memberService.NewMemberService(memberRepositroy,db,validate)
	memberController := memberController.NewMemberController(memberService)

	bookRepository := bookRepository.NewBookRepository()
	bookService := bookService.NewBookService(bookRepository,db,validate)
	bookController := bookController.NewBookController(bookService)

	transactionRepository := transactionRepository.NewTransactionRepository()
	transactionService := transactionService.NewTransactionService(transactionRepository,memberRepositroy,bookRepository,db,validate)
	transactionController := transactionController.NewTransactionController(transactionService)

	router := httprouter.New()

	//member
	router.GET("/library/members",memberController.FindAllMember)
	router.GET("/library/members/:memberId",memberController.FindMemberById)
	router.POST("/library/members",memberController.AddMember)
	router.DELETE("/library/members/:memberId",memberController.DeleteMember)

	//book
	router.GET("/library/books",bookController.FindAllBook)
	router.GET("/library/books/:bookId",bookController.FindBookById)
	router.POST("/library/books",bookController.AddBook)
	router.PATCH("/library/books/:bookId",bookController.UpdateBook)
	router.DELETE("/library/books/:bookId",bookController.DeleteBook)

	//transaction
	router.GET("/library/findborrowed/:bookId",transactionController.FindMemberWhoBorrowBook)
	router.POST("/library/borrow/:memberId",transactionController.BorrowBook)
	router.PATCH("/library/return/:memberId",transactionController.ReturnBook)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}