package Web

import "time"

type TransactionFindResponse struct {
	MemberId    string
	MemberName  string
	MemberEmail string
	BookId      string
	BookTitle   string
	LoanDate    time.Time
	IsReturn string
}