package main

import (
	"github.com/krishh1at/simple_app/db"

	"github.com/krishh1at/simple_app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Migration()

	gin.ForceConsoleColor()
	router := gin.Default()
	routes.InitRouter(router)
	router.Run(":3000")
}
