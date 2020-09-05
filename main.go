package main

import (
	"fmt"

	"github.com/krishh1at/simple_app/config"
	"github.com/krishh1at/simple_app/models"
	"github.com/krishh1at/simple_app/routers"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var err error
	config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}

	config.DB.AutoMigrate(&models.User{})

	gin.ForceConsoleColor()
	router := gin.Default()
	routers.InitRouter(router)
	router.Run(":8080")
}
