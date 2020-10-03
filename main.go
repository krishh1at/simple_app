package main

import (
	"fmt"

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
	router.Use(sessions.Sessions("mysession", store))

	routes.InitRouter(router)

	s := fmt.Sprintf("/users/%v", 10)
	fmt.Printf("%v : %T", s, s)
	router.Run(":8080")
}
