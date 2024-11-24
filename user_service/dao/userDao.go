package dao

import (
	"fmt"
	models "user_service/models"
	"gorm.io/gorm"
)


type UserDao struct {
    DBTemplate *gorm.DB
}

func (userDao * UserDao) CreateUser(user *models.SysUser) error {
	if userDao.DBTemplate == nil {
        return fmt.Errorf("database connection is nil")
    }

    // 调用 Save 保存用户
    result := userDao.DBTemplate.Create(user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// 修改用户信息
func (userDao * UserDao) UpdateUser(user *models.SysUser) error {
	if userDao.DBTemplate == nil {
    	return fmt.Errorf("database connection is nil")
    }

    // 调用 Save 保存用户
    result := userDao.DBTemplate.Create(user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}


// 根据邮箱获取用户
func (userDao * UserDao) GetUserByEmail(email string) (models.SysUser, error) {
    var user models.SysUser
    result := userDao.DBTemplate.Where("email = ?", email).First(&user)
    if result.Error != nil {
        return user, result.Error
    }
    return user, nil
}

// 根据手机号获取用户
func (userDao * UserDao) GetUserByPhone(phone string) (models.SysUser, error) {
    var user models.SysUser
    result := userDao.DBTemplate.Where("phone = ?", phone).First(&user)
    if result.Error != nil {
        return user, result.Error
    }
    return user, nil
}