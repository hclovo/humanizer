package controller

import (
	"log"
	"net/http"
	"user_service/request"
	"user_service/service"

	"github.com/gin-gonic/gin"
)


type UserController interface {
	RegisterUser(ctx *gin.Context) 
	GetUser(ctx *gin.Context) 
}

func RegisterUser(ctx *gin.Context) {
	log.Println("RegisterUser ...")
	var userRequest request.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
	    ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	status, result := service.RegisterUser(&userRequest.User)
	ctx.JSON(status, result)
}

func GetEmailCode(ctx *gin.Context) {
	log.Println("GetEmailCode ...")
	var request request.EmailCodeRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	status, result := service.GetEmailCode(&request)
	ctx.JSON(status, result)
}


func LoginUser(ctx *gin.Context) {
	log.Println("LoginUser ...")
	var userRequest request.UserLoginRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
	    ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if userRequest.Type == 0 {
		status, result := service.LoginBySupabase(&userRequest.User)
		ctx.JSON(status, result)
	}
	
}

func GetPhoneCode(ctx *gin.Context) {
	log.Println("GetPhoneCode ...")
	var request request.PhoneCodeRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	status, result := service.GetPhoneCode(request.Phone)
	ctx.JSON(status, result)
}