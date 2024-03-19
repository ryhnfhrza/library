package repository

import (
	"context"
	"database/sql"
	"errors"
	domain "library/Model/Domain"
	"library/helper"
)

type TransactionRepositoryImpl struct{

}

func NewTransactionRepository()TransactionRepository{
	return &TransactionRepositoryImpl{}
}

func(repository *TransactionRepositoryImpl)BorrowBook(ctx context.Context,tx *sql.Tx,transaction domain.Transaction)domain.Transaction{
 SQL := "insert into book_transaction (member_id, book_id, loan_date, is_return) values (?, ?, ?, ?);"
 _,err := tx.ExecContext(ctx,SQL,transaction.MemberId,transaction.BookId,transaction.LoanDate,transaction.IsReturn);
 helper.PanicIfError(err)

 return transaction
}

func(repository *TransactionRepositoryImpl)ReturnBook(ctx context.Context,tx *sql.Tx,transaction domain.Transaction)domain.Transaction{
	SQL := `update book_transaction set is_return = ?,loan_date = ? where member_id = ? and book_id = ? `
	_,err := tx.ExecContext(ctx,SQL,transaction.IsReturn,transaction.LoanDate,transaction.MemberId,transaction.BookId)
	helper.PanicIfError(err)

	return transaction
}

func(repository *TransactionRepositoryImpl)FindMemberWhoBorrowBook(ctx context.Context,tx *sql.Tx,bookId string)(domain.TransactionTracking,error){
	SQL := "select m.id,m.name,m.email,b.id,b.title,bt.loan_date,bt.is_return from book_transaction bt inner join member m on bt.member_id = m.id inner join book b on bt.book_id = b.id where b.id = ?"
	
	rows,err := tx.QueryContext(ctx,SQL,bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	transaction := domain.TransactionTracking{}
	if rows.Next(){
		err := rows.Scan(&transaction.MemberId,&transaction.MemberName,&transaction.MemberEmail,&transaction.BookId,&transaction.BookTitle,&transaction.LoanDate,&transaction.IsReturn)
		helper.PanicIfError(err)
		return transaction,nil
	}else{
		return transaction,errors.New("Book id " + bookId + " Available")
	}
}

func(repository *TransactionRepositoryImpl)CheckBookAvailable(ctx context.Context,tx *sql.Tx,bookId string)(domain.Transaction){
	SQL := "select  is_return from book_transaction where book_id = ? order by loan_date desc limit 1"

	rows,err := tx.QueryContext(ctx,SQL,bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	transaction := domain.Transaction{}
	if rows.Next(){
		err := rows.Scan(&transaction.IsReturn)
		helper.PanicIfError(err)
		return transaction
	}else{
		return transaction
	}
}
func(repository *TransactionRepositoryImpl)CheckBookInTransaction(ctx context.Context,tx *sql.Tx,bookId string)bool{
	SQL := `select book_id from book_transaction where book_id = ?`
	rows,err := tx.QueryContext(ctx,SQL,bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next(){
		return true
	}else{
		return false
	}
}

func(repository *TransactionRepositoryImpl)CheckBeforeReturning(ctx context.Context,tx *sql.Tx,bookId ,memberId string) (bool,error){
	SQL := `select member_id,book_id from book_transaction where member_id = ? and book_id = ? and is_return = "no"`
	rows,err := tx.QueryContext(ctx,SQL,memberId,bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next(){
		return true,nil
	}else{
		return false,errors.New("data not found")
	}
}

func(repository *TransactionRepositoryImpl)CheckIsReturn(ctx context.Context,tx *sql.Tx,bookId string)(bool,error){
	SQL := `select book_id from book_transaction where book_id = ? and is_return = "no"`
	rows,err := tx.QueryContext(ctx,SQL,bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next(){
		return true,nil
	}else{
		return false,errors.New("the book is not being borrowed by anyone")
	}
}


