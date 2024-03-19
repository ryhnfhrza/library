package Web

type TransactionBorrowAndReturnRequest struct {
	MemberId string `validate:"required,max=8,min=8" json:"member_id"`
	BookId   string `validate:"required,max=10,min=1" json:"book_id"`
	IsReturn string `json:"is_return"`
}