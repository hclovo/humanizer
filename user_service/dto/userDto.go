package dto

import (
	"user_service/models"
)


type UserDto struct {
    *models.SysUser
	ComfirmPassword string `json:"confirm_password"`
	Code string `json:"code"`
}