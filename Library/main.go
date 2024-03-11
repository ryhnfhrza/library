package main

import (
	"library/app"
	bookController "library/controller/BookController"
	memberController "library/controller/MemberController"
	"library/exception"
	"library/helper"
	bookRepository "library/repository/BookRepository"
	memberRepository "library/repository/MemberRepository"
	bookService "library/service/BookService"
	memberService "library/service/MemberService"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main(){

	db := app.NewDb()
	validate := validator.New()

	memberRepositroy := memberRepository.NewMemberRepository()
	memberService := memberService.NewMemberService(memberRepositroy,db,validate)
	memberController := memberController.NewMemberController(memberService)

	bookRepository := bookRepository.NewBookRepository()
	bookService := bookService.NewBookService(bookRepository,db,validate)
	bookController := bookController.NewBookController(bookService)

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
	router.PUT("/library/books/:bookId",bookController.UpdateBook)
	router.DELETE("/library/books/:bookId",bookController.DeleteBook)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}