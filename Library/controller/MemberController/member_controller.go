package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MemberController interface{
	AddMember(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	DeleteMember(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	FindMemberById(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
	FindAllMember(writer http.ResponseWriter,request *http.Request,params httprouter.Params)
}