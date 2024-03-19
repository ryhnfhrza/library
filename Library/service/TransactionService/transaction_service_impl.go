package service

import (
	"context"
	"database/sql"
	domain "library/Model/Domain"
	Web "library/Model/Web/TransactionWeb"
	"library/exception"
	"library/helper"
	repositoryBook "library/repository/BookRepository"
	repositoryMember "library/repository/MemberRepository"
	repositoryTransaction "library/repository/TransactionRepository"
	"time"

	"github.com/go-playground/validator/v10"
)

type TransactionServiceImpl struct{
	TransactionRepository repositoryTransaction.TransactionRepository
	MemberRepository repositoryMember.MemberRepository 
	BookRepository repositoryBook.BookRepository
	Db *sql.DB
	Validate *validator.Validate
}

func NewTransactionService(transactionRepository repositoryTransaction.TransactionRepository,memberRepository repositoryMember.MemberRepository,bookRepository repositoryBook.BookRepository,db *sql.DB,validate *validator.Validate)TransactionService{
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		MemberRepository: memberRepository,
		BookRepository: bookRepository,
		Db: db,
		Validate: validate,
	}
}

func(service *TransactionServiceImpl)BorrowBook(ctx context.Context, request Web.TransactionBorrowAndReturnRequest)Web.TransactionBorrowReturnResponse{
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//validate member
	member,err := service.MemberRepository.FindMemberById(ctx,tx,request.MemberId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	//validate Book available or not
	book,err := service.BookRepository.FindBookById(ctx,tx,request.BookId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	//check book already in transaction table or not
	resultBook := service.TransactionRepository.CheckBookInTransaction(ctx,tx,book.Id)
	if resultBook {
		
		//validate whether the book is being borrowed or not
		result := service.TransactionRepository.CheckBookAvailable(ctx,tx,book.Id)
		if result.IsReturn == "no"{
			err := exception.NewConflictError("Cannot borrow book " + book.Id + " because the book is being borrowed")
			panic(err)
		}
		
	}

	//get time right now
	curentTime := time.Now()

	transaction := domain.Transaction{
		MemberId: member.Id,
		BookId: request.BookId,
		LoanDate: curentTime,
		IsReturn: "no",
	}

	transaction = service.TransactionRepository.BorrowBook(ctx,tx,transaction)

	return helper.ToTransactionBorrowReturnResponse(transaction)
}

func(service *TransactionServiceImpl)ReturnBook(ctx context.Context,request Web.TransactionBorrowAndReturnRequest)Web.TransactionBorrowReturnResponse{
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
			result,err := service.TransactionRepository.CheckBeforeReturning(ctx,tx,request.BookId,request.MemberId)
			if err != nil{
					panic(exception.NewNotFoundError(err.Error()))
				 }
			if !result{
				panic("Data not found")
			}

			request.IsReturn = "yes"
			
			//get current time
			curentTime := time.Now()

	transaction := domain.Transaction{}

	transaction.MemberId = request.MemberId
	transaction.BookId = request.BookId
	transaction.LoanDate = curentTime
	transaction.IsReturn = request.IsReturn

	transaction = service.TransactionRepository.ReturnBook(ctx,tx,transaction)

	return helper.ToTransactionBorrowReturnResponse(transaction)
}

func(service *TransactionServiceImpl)FindMemberWhoBorrowBook(ctx context.Context,bookId string)Web.TransactionFindResponse{
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book,err := service.BookRepository.FindBookById(ctx,tx,bookId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	result,err := service.TransactionRepository.CheckIsReturn(ctx,tx,book.Id)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	 }
	if !result{
		panic("Data not found")
	}

	transaction,err := service.TransactionRepository.FindMemberWhoBorrowBook(ctx,tx,book.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTransactionFindResponse(transaction)
}

