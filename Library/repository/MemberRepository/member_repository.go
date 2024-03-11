package repository

import (
	"context"
	"database/sql"
	domain "library/Model/Domain"
)

type MemberRepository interface {
	AddMember(ctx context.Context, tx *sql.Tx,member domain.Member)domain.Member
	DeleteMember(ctx context.Context, tx *sql.Tx,member domain.Member)
	FindMemberById(ctx context.Context, tx *sql.Tx,memberId string)(domain.Member,error)
	FindAllMember(ctx context.Context, tx *sql.Tx)[]domain.Member
}