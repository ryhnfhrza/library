package repository

import (
	"context"
	"database/sql"
	"errors"
	domain "library/Model/Domain"
	"library/helper"
)

type MemberRepositoryImpl struct{

}

func NewMemberRepository()MemberRepository{
	return &MemberRepositoryImpl{}
}

func(repository *MemberRepositoryImpl)AddMember(ctx context.Context, tx *sql.Tx,member domain.Member)domain.Member{
	SQL := "insert into member (id,name,birth_date,email,address) values(SUBSTRING(UUID(), 1, 8),?,?,?,?)"
	_,err := tx.ExecContext(ctx,SQL,member.Name,member.BirthDate,member.Email,member.Address)
	helper.PanicIfError(err)

	return member
}

func(repository *MemberRepositoryImpl)DeleteMember(ctx context.Context, tx *sql.Tx,member domain.Member){
	SQL := "Delete from member where id = ? "
	_,err := tx.ExecContext(ctx,SQL,member.Id)
	helper.PanicIfError(err)

}

func(repository *MemberRepositoryImpl)FindMemberById(ctx context.Context, tx *sql.Tx,memberId string)(domain.Member,error){
	SQL := "select id,name,email,birth_date,address from member where id = ?"
	rows,err := tx.QueryContext(ctx,SQL,memberId)
	helper.PanicIfError(err)
	defer rows.Close()

	member := domain.Member{}
	if rows.Next(){
		err := rows.Scan(&member.Id,&member.Name,&member.Email,&member.BirthDate,&member.Address)
		helper.PanicIfError(err)
		return member,nil

	}else{
		return member,errors.New("member with id "+memberId +" not found")
	}
}

func(repository *MemberRepositoryImpl)FindAllMember(ctx context.Context, tx *sql.Tx)[]domain.Member{
	SQL := "select id,name,email,birth_date,address from member"
	rows,err:=tx.QueryContext(ctx,SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var members []domain.Member
	for rows.Next(){
		member := domain.Member{}
		err := rows.Scan(&member.Id,&member.Name,&member.Email,&member.BirthDate,&member.Address)
		helper.PanicIfError(err)
		members = append(members,member)
	}
	return members
}
