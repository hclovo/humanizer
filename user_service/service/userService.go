package service

import (
	"log"
	"net/http"
	"reflect"
	"time"
	"user_service/common"
	"user_service/dao"
	"user_service/dto"
	"user_service/models"
	"user_service/request"
	"user_service/utils"

	"github.com/gin-gonic/gin"
	"github.com/supabase-community/gotrue-go/types"
	"golang.org/x/crypto/bcrypt"
)

var db = common.GetDB()

var dbTemplate = &dao.UserDao{DBTemplate: db}

var redisclient= common.GetRedisClient()

var redisTemplate = &common.RedisTemplate{RedisClient: redisclient}

var supabaseClient = common.GetSupabaseClient()


func checkUserbyEmail(email string) bool {
	user, _ := dbTemplate.GetUserByEmail(email)
	var zeorUser models.SysUser
	return !reflect.DeepEqual(user, zeorUser)
}

func checkUser(user *models.SysUser) bool {
    var zeorUser models.SysUser
    return !reflect.DeepEqual(user, zeorUser)
}

func RegisterUser(user *dto.UserDto) (int, *gin.H) {
    email := user.Email
    if checkUserbyEmail(email) {
        return http.StatusBadRequest, &gin.H{"error": "邮箱已存在"}
    }

	userId, err := utils.GetUID()

	if err != nil {
		log.Println("Error generating user ID: ?", err.Error())
		return http.StatusInternalServerError, &gin.H{"error": "系统报错，请联系管理员"}
	}

    password := user.Password
	comfirmPassword := user.ComfirmPassword

    // code := user.Code

	// verfiy_code, err:= redisTemplate.GetValue(email)

	// if err != nil {
	//     return http.StatusInternalServerError, &gin.H{"error": "系统报错，请联系管理员，" + err.Error()}
	// }

    // if code != verfiy_code {
    //     return http.StatusBadRequest, &gin.H{"error": "邮箱验证码输入错误，请重试"}
    // }

	if password != comfirmPassword {
	    return http.StatusBadRequest, &gin.H{"error": "两次密码输入不一致，请重新输入"}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return http.StatusInternalServerError, &gin.H{"error": "密码加密失败，请联系管理员"}
    }
	
	supabaseRequest := types.SignupRequest{
		Email: email,
		Password: password,
	}

	// 注册到supabase
	_, err = supabaseClient.Auth.Signup(supabaseRequest)
	if err != nil {
		log.Println("Error registering user: ?", err.Error())
		return http.StatusInternalServerError, &gin.H{"error": "注册失败，请联系管理员"}
	}


    registUser := models.SysUser{
		UserId: userId,
		Email:   email,
		Password: string(hashedPassword),
		Nickname: user.Nickname,
		Avatar:  user.Avatar,
		Phone:   user.Phone,
		Username: user.Username,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		IsDeleted: 0,
	}

	err = dbTemplate.CreateUser(&registUser)

	if err != nil {
	    return http.StatusInternalServerError, &gin.H{"error": "系统报错，请联系管理员，" + err.Error()}
	}

	return http.StatusOK, &gin.H{"message": "注册成功"}
}


func GetEmailCode(request *request.EmailCodeRequest) (int, *gin.H) {
	email := request.Email
	
	// 生成验证码
	code := utils.RandomString(6)
	log.Println("code: ", code)
	// 发送验证码
	err := utils.SendEmailCode(email, code)

	if err != nil {
	    return http.StatusInternalServerError, &gin.H{"error": "系统报错，请联系管理员，" + err.Error()}
	}

	// 存储验证码到数据库
	err = redisTemplate.SetValue(email, code, 5 * time.Minute)
	if err != nil {
	    return http.StatusInternalServerError, &gin.H{"error": "系统报错，请联系管理员，" + err.Error()}
	}
	return http.StatusOK, &gin.H{"message": "验证码发送成功, 请查收"}
}

func LoginByPhone(user *dto.UserDto) (int, *gin.H) {
	userId := user.UserId
	dbUser, _ := dbTemplate.GetUserByPhone(user.Phone)
	code := user.Code
	verfiy_code, err := redisTemplate.GetValue(user.Phone)
	if err != nil {
	    return http.StatusBadRequest, &gin.H{"error": "验证码失效"}
	}
	if code != verfiy_code {
	    return http.StatusBadRequest, &gin.H{"error": "验证码错误"}
	}
	if !checkUser(&dbUser) {
		// 用户不存在就注册用户
		userId, err = utils.GetUID()
		if err != nil {
			return http.StatusInternalServerError, &gin.H{"error": "系统报错，请联系管理员"}
		}
		registUser := models.SysUser{
			UserId: userId,
			Email:   user.Email,
			Password: user.Password,
			Nickname: user.Nickname,
			Avatar:  user.Avatar,
			Phone:   user.Phone,
			Username: user.Username,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			IsDeleted: 0,
		}
	
		err = dbTemplate.CreateUser(&registUser)
	
		if err != nil {
			return http.StatusInternalServerError, &gin.H{"error": "系统报错，请联系管理员，" + err.Error()}
		}
	}

	token, err := utils.GenerateJWT(userId)
	if err != nil {
		return http.StatusInternalServerError, &gin.H{"error": "Token generation failed"}
	}

	return http.StatusOK, &gin.H{"message": "登录成功", "token": token}
}


func CheckPassword(password string, hashedPassword string) bool {
	bryotPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return false
    }
	if string(bryotPassword) == hashedPassword {
	    return true
	}
	return false
}


func LoginByEmail(user *dto.UserDto) (int, *gin.H) {
	dbUser, _ := dbTemplate.GetUserByEmail(user.Email)
	if checkUser(&dbUser) {
		return http.StatusBadRequest, &gin.H{"error": "用户不存在"}
	}
	if !CheckPassword(user.Password, dbUser.Password) {
		return http.StatusBadRequest, &gin.H{"error": "密码错误"}
	}
	userId := dbUser.UserId
	token, err := utils.GenerateJWT(userId)
	if err != nil {
		return http.StatusInternalServerError, &gin.H{"error": "Token generation failed"}
	}

	return http.StatusOK, &gin.H{"message": "登录成功", "token": token}
}
func LoginUser(user *dto.UserDto) (int, *gin.H) {
	if (user.Phone != "") {
		return LoginByPhone(user)
	} 
	
	return http.StatusOK, &gin.H{"message": "登录成功"}
}

func GetPhoneCode(phone string) (int, *gin.H) {
	// TODO: 调用短信服务发送验证码
	return http.StatusOK, &gin.H{"message": "验证码发送成功"}
}

func LoginBySupabase(user *dto.UserDto) (int, *gin.H) {
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    // if err != nil {
    //     return http.StatusInternalServerError, &gin.H{"error": "密码加密失败，请联系管理员"}
    // }
	email := user.Email
	log.Println("email:", email)
	password := user.Password
	log.Println("password:", password)

	tokenResponse, err := supabaseClient.Auth.SignInWithEmailPassword(email, password)
	log.Println("token response:", tokenResponse)
	if err != nil {
		log.Println("登录失败：", err.Error())
		return http.StatusInternalServerError, &gin.H{"error": "登录失败，请联系管理员"}
	}
	return http.StatusOK, &gin.H{"message": "登录成功", "token": tokenResponse.AccessToken}
}