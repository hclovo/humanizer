package routers

import (
	"net/http"
	"user_service/controller"
	"user_service/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/register", controller.RegisterUser)
	r.POST("/login", controller.LoginUser)
	r.POST("/get_email_code", controller.GetEmailCode)


	auth := r.Group("/auth", utils.AuthMiddleware)
	auth.GET("/profile", profile)
	return r
}


func profile(c *gin.Context) {
	username := c.GetString("username")
	c.JSON(http.StatusOK, gin.H{"message": "Welcome " + username})
}
