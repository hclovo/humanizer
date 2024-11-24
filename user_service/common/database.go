package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"user_service/utils"
)

var DB *gorm.DB

func InitDatabase() *gorm.DB{
	config := utils.AppConfig.Database
	log.Printf("Connecting to database: %s\n", config.DSN)
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Database connected successfully!")
	DB = db
	return db
}

func GetDB() *gorm.DB {
	if DB == nil {
		utils.LoadConfig()
		InitDatabase()
	}

    return DB
}
