package controller

import (
	web "library/Model/Web"
	Web "library/Model/Web/MemberWeb"
	"library/helper"
	memberService "library/service/MemberService"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MemberControllerImpl struct{
	MemberService memberService.MemberService
}

func NewMemberController(memberService memberService.MemberService)MemberController{
	return &MemberControllerImpl{
		MemberService: memberService,
	}
}

func(controller *MemberControllerImpl)AddMember(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	memberAddRequest := Web.MemberAddRequest{}
	helper.ReadFromRequestBody(request,&memberAddRequest)

	memberResponse := controller.MemberService.AddMember(request.Context(),memberAddRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: memberResponse,
	}
	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *MemberControllerImpl)DeleteMember(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	memberId := params.ByName("memberId")

	controller.MemberService.DeleteMember(request.Context(),memberId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *MemberControllerImpl)FindMemberById(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	memberId := params.ByName("memberId")

	memberResponse := controller.MemberService.FindMemberById(request.Context(),memberId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: memberResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *MemberControllerImpl)FindAllMember(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	memberResponse := controller.MemberService.FindAllMember(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: memberResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}
