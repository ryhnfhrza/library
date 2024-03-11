package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BookController interface{
	AddBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params)
	UpdateBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params)
	DeleteBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params)
	FindBookById(writer http.ResponseWriter, request *http.Request,params httprouter.Params)
	FindAllBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params)
}