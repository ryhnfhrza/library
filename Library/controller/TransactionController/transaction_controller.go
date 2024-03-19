package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TransactionController interface{
	BorrowBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params)
	ReturnBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params)
	FindMemberWhoBorrowBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params)
}