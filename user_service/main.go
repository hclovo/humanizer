package main

import (
	"log"
	// "user_service/common"
	routers "user_service/routers"
	config "user_service/utils"

	"github.com/gin-gonic/gin"
)

// func _init() {
// 	log.Println("init ...")
// 	config.LoadConfig()
// 	db := common.InitDatabase()
// 	_, err := db.DB()
// 	if err != nil {
// 	    panic("failed to connect to database!" + err.Error())
// 	}

// 	_ = common.InitRedis()
// }

func main() {
	config.LoadConfig()
	log.Println("main ...")
	// _init()

	port := config.AppConfig.Server.Port
	r := gin.Default()
	r = routers.RegisterRoutes(r)
	panic(r.Run(":" + port))

}
