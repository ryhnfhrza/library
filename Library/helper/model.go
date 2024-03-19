package helper

import (
	domain "library/Model/Domain"
	bookResponse "library/Model/Web/BookWeb"
	memberResponse "library/Model/Web/MemberWeb"
	transactionResponse "library/Model/Web/TransactionWeb"
)

func ToMemberResponse(member domain.Member)memberResponse.MemberResponse{
	return memberResponse.MemberResponse{
		Id: member.Id,
		Name: member.Name,
		Email: member.Email,
		BirthDate: member.BirthDate,
		Address: member.Address,
	}

}

func ToMemberResponses(members []domain.Member)[]memberResponse.MemberResponse{
	var memberResponses []memberResponse.MemberResponse
	for _,member := range members{
		memberResponses = append(memberResponses, ToMemberResponse(member))
	}
	return memberResponses
}

func ToBookResponse(book domain.Book)bookResponse.BookResponse{
	return bookResponse.BookResponse{
		Id: book.Id,
		Title: book.Title,
		Author: book.Author,
		Category: book.Category,
	}
}

func ToBookResponses(books []domain.Book)[]bookResponse.BookResponse{
	var bookResponses []bookResponse.BookResponse
	for _,book := range books{
		bookResponses = append(bookResponses, ToBookResponse(book))
	}
	return bookResponses
}

func ToTransactionBorrowReturnResponse(transaction domain.Transaction)transactionResponse.TransactionBorrowReturnResponse{
	return transactionResponse.TransactionBorrowReturnResponse{
		MemberId: transaction.MemberId,
		BookId: transaction.BookId,
		LoanDate: transaction.LoanDate,
		IsReturn: transaction.IsReturn,
	}
}

func ToTransactionFindResponse(transaction domain.TransactionTracking)transactionResponse.TransactionFindResponse{
	return transactionResponse.TransactionFindResponse{
		MemberId: transaction.MemberId,
		MemberName: transaction.MemberName,
		MemberEmail: transaction.MemberEmail,
		BookId: transaction.BookId,
		BookTitle: transaction.BookTitle,
		LoanDate: transaction.LoanDate,
		IsReturn: transaction.IsReturn,
	}
}