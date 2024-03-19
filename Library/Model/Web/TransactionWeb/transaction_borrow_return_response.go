package Web

import (
	"time"
)

type TransactionBorrowReturnResponse struct{
	MemberId string `json:"member_id"`
	BookId   string `json:"book_id"`
	LoanDate time.Time `json:"loan_date"`
	IsReturn string `json:"is_return"`
}