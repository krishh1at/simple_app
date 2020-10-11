package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/krishh1at/simple_app/db"
	"github.com/krishh1at/simple_app/routes"
)

func main() {
	db.Migration()

	gin.ForceConsoleColor()
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("simple_app", store))

	routes.InitRouter(router)
	router.Run(":8080")
}
