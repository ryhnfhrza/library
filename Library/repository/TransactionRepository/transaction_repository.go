package repository

import (
	"context"
	"database/sql"
	domain "library/Model/Domain"
)

type TransactionRepository interface {
	BorrowBook(ctx context.Context,tx *sql.Tx,transaction domain.Transaction)domain.Transaction
	ReturnBook(ctx context.Context,tx *sql.Tx,transaction domain.Transaction)domain.Transaction
	FindMemberWhoBorrowBook(ctx context.Context,tx *sql.Tx,bookId string)(domain.TransactionTracking,error)
	CheckBookAvailable(ctx context.Context,tx *sql.Tx,bookId string)(domain.Transaction)
	CheckBookInTransaction(ctx context.Context,tx *sql.Tx,bookId string)bool
	CheckBeforeReturning(ctx context.Context,tx *sql.Tx,bookId ,memberId string) (bool,error)
	CheckIsReturn(ctx context.Context,tx *sql.Tx,bookId string)(bool,error)

}