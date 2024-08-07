package controller

import (
	"net/http"

	"github.com/ryhnfhrza/Golang-To-Do-List-API/helper"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/web"
	service "github.com/ryhnfhrza/Golang-To-Do-List-API/service/AuthService"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/util"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authservice service.AuthService)AuthController{
	return &AuthControllerImpl{
		AuthService: authservice,
	}
}

func(controller *AuthControllerImpl)Registration(writer http.ResponseWriter,request *http.Request){
	RegistrationRequest := web.RegistrationRequest{}
	helper.ReadFromRequestBody(request,&RegistrationRequest)

	RegistrationResponse := controller.AuthService.Registration(request.Context(),RegistrationRequest)
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Data: RegistrationResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request) {
	LoginRequest := web.LoginRequest{}
	helper.ReadFromRequestBody(request, &LoginRequest)

	LoginResponse, tokenAlgo, err := controller.AuthService.Login(request.Context(), LoginRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	token, err := tokenAlgo.SignedString([]byte(util.JWT_KEY))
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Set cookie with token
	http.SetCookie(writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   LoginResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}


func(controller *AuthControllerImpl)Logout(writer http.ResponseWriter,request *http.Request){
	http.SetCookie(writer ,& http.Cookie{
		Name: "token",
		Path: "/",
		Value: "",
		HttpOnly: true,
		MaxAge: -1,
	})
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Ok",
		Data: "Logout",
	}

	helper.WriteToResponseBody(writer,webResponse)

}

