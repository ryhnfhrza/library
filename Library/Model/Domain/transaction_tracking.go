package domain

import "time"

type TransactionTracking struct {
	MemberId    string
	MemberName  string
	MemberEmail string
	BookId      string
	BookTitle   string
	LoanDate    time.Time
	IsReturn string
}