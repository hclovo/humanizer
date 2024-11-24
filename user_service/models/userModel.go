package models

import (
	"fmt"
	"os"
	"time"
	"user_service/common"
)


type SysUser struct {
	UserId       string `json:"user_id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Nickname     string `json:"nickname"`
	Avatar       string `json:"avatar"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
	IsDeleted    int    `json:"is_deleted"`
}

var db = common.GetDB()

func (user * SysUser) TableName() string {
	fmt.Fprintln(os.Stdout, user.Username)
	return "sys_user"
}

// 创建用户
func (user * SysUser) CreateUser() error {
	if db == nil {
        return fmt.Errorf("database connection is nil")
    }

    if user == nil {
        return fmt.Errorf("SysUser instance is nil")
    }

    // 调用 Save 保存用户
    result := db.Create(user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// 修改用户信息
func (user * SysUser) UpdateUser() error {
	if db == nil {
        return fmt.Errorf("database connection is nil")
    }

    if user == nil {
        return fmt.Errorf("SysUser instance is nil")
    }

    // 调用 Save 保存用户
    result := db.Create(user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}


// 根据邮箱获取用户
func (user * SysUser) GetUserByEmail(email string) (SysUser, error) {
    var u SysUser
    result := db.Where("email = ?", email).First(&u)
    if result.Error != nil {
        return u, result.Error
    }
    return u, nil
}       