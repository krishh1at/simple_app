package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/krishh1at/simple_app/app/models"
	"github.com/krishh1at/simple_app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migration() {
	var err error
	config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
	}

	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Post{})
}
