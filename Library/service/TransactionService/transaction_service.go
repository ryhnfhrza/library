package service

import (
	"context"
	Web "library/Model/Web/TransactionWeb"
)

type TransactionService interface {
	BorrowBook(ctx context.Context, request Web.TransactionBorrowAndReturnRequest)Web.TransactionBorrowReturnResponse
	ReturnBook(ctx context.Context,request Web.TransactionBorrowAndReturnRequest)Web.TransactionBorrowReturnResponse
	FindMemberWhoBorrowBook(ctx context.Context,bookId string)Web.TransactionFindResponse
}