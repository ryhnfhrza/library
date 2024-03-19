package controller

import (
	web "library/Model/Web"
	Web "library/Model/Web/TransactionWeb"
	"library/helper"
	service "library/service/TransactionService"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TransactionControllerImpl struct{
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService)TransactionController{
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

func(controller *TransactionControllerImpl)BorrowBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	transactionBorrowRequest := Web.TransactionBorrowAndReturnRequest{}
	helper.ReadFromRequestBody(request,&transactionBorrowRequest)

	memberId := params.ByName("memberId")
	transactionBorrowRequest.MemberId = memberId

	transactionResponse := controller.TransactionService.BorrowBook(request.Context(),transactionBorrowRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: transactionResponse,
	}
	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *TransactionControllerImpl)ReturnBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	transactionReturnRequest := Web.TransactionBorrowAndReturnRequest{}
	helper.ReadFromRequestBody(request,&transactionReturnRequest)

	memberId := params.ByName("memberId")
	transactionReturnRequest.MemberId = memberId

	transactionResponse := controller.TransactionService.ReturnBook(request.Context(),transactionReturnRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: transactionResponse,
	}
	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *TransactionControllerImpl)FindMemberWhoBorrowBook(writer http.ResponseWriter, request *http.Request,params httprouter.Params){
	bookId := params.ByName("bookId")

	transactionResponse := controller.TransactionService.FindMemberWhoBorrowBook(request.Context(),bookId)
	webResponse:= web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: transactionResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}
