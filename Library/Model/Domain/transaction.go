package domain

import "time"

type Transaction struct {
	MemberId string
	BookId   string
	LoanDate time.Time
	IsReturn string
}