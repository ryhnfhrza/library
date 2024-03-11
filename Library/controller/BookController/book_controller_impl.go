package controller

import (
	web "library/Model/Web"
	bookWeb "library/Model/Web/BookWeb"
	"library/helper"
	bookService "library/service/BookService"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BookControllerImpl struct{
	BookService bookService.BookService
}

func NewBookController(bookService bookService.BookService)BookController{
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func(controller *BookControllerImpl)AddBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	bookAddRequest := bookWeb.BookAddRequest{}
	helper.ReadFromRequestBody(request,&bookAddRequest)

	bookResponse := controller.BookService.AddBook(request.Context(),bookAddRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bookResponse,
	}
	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *BookControllerImpl)UpdateBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	bookUpdateRequest := bookWeb.BookUpdateRequest{}
	helper.ReadFromRequestBody(request,&bookUpdateRequest)

	bookId := params.ByName("bookId")
	bookUpdateRequest.Id = bookId

	bookResponse := controller.BookService.UpdateBook(request.Context(),bookUpdateRequest)
	webResponse:= web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bookResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *BookControllerImpl)DeleteBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	bookId := params.ByName("bookId")

	controller.BookService.DeleteBook(request.Context(),bookId)
	webResponse:= web.WebResponse{
		Code: 200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *BookControllerImpl)FindBookById(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	bookId := params.ByName("bookId")

	bookResponse := controller.BookService.FindBookById(request.Context(),bookId)
	webResponse:= web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bookResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *BookControllerImpl)FindAllBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	bookResponses := controller.BookService.FindAllBook(request.Context())
	webResponse:= web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: bookResponses,
	}

	helper.WriteToResponseBody(writer,webResponse)
}
