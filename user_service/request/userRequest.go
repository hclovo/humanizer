package request

import (
	"user_service/dto"
)

type UserLoginRequest struct {
    User dto.UserDto `json:"user"`
	Type int `json:"type" comment:"是否是微信登录 0不是,1是,默认0"`
}

type UserRegisterRequest struct {
    User dto.UserDto `json:"user"`
}

type EmailCodeRequest struct {
    Email string 	`json:"email"`
}

type PhoneCodeRequest struct {
    Phone string 	`json:"phone"`
}