package service

import (
	"context"
	Web "library/Model/Web/MemberWeb"
)

type MemberService interface{
	AddMember(ctx context.Context,request Web.MemberAddRequest) Web.MemberResponse
	DeleteMember(ctx context.Context,memberId string)
	FindMemberById(ctx context.Context,memberId string) Web.MemberResponse
	FindAllMember(ctx context.Context)[]Web.MemberResponse
}