package service

import (
	"context"
	"database/sql"
	domain "library/Model/Domain"
	web "library/Model/Web/MemberWeb"
	"library/exception"
	"library/helper"
	repository "library/repository/MemberRepository"
	"time"

	"github.com/go-playground/validator/v10"
)

type MemberServiceImpl struct{
	MemberRepository repository.MemberRepository
	Db *sql.DB
	Validate *validator.Validate
}

func NewMemberService(memberRepository repository.MemberRepository ,db *sql.DB,validate *validator.Validate)MemberService{
	return &MemberServiceImpl{
		MemberRepository: memberRepository,
		Db: db,
		Validate: validate,
	}
}

func(service *MemberServiceImpl)AddMember(ctx context.Context,request web.MemberAddRequest) web.MemberResponse{
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	formatter := "2006-01-02"
	birthDate := request.BirthDate
	parseTime , err := time.Parse(formatter,birthDate)
	helper.PanicIfError(err)


	member := domain.Member{
		Name: request.Name,
		Email: request.Email,
		BirthDate: parseTime,
		Address: request.Address,
	}
	member = service.MemberRepository.AddMember(ctx,tx,member)

	return helper.ToMemberResponse(member)
}

func(service *MemberServiceImpl)DeleteMember(ctx context.Context,memberId string){
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	member,err := service.MemberRepository.FindMemberById(ctx,tx,memberId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.MemberRepository.DeleteMember(ctx,tx,member)
}

func(service *MemberServiceImpl)FindMemberById(ctx context.Context,memberId string) web.MemberResponse{
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	member,err := service.MemberRepository.FindMemberById(ctx,tx,memberId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMemberResponse(member)
}

func(service *MemberServiceImpl)FindAllMember(ctx context.Context)[]web.MemberResponse{
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	members := service.MemberRepository.FindAllMember(ctx,tx)

	return helper.ToMemberResponses(members)
}
