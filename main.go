package main

import (
	"fmt"

	"github.com/krishh1at/simple_app/db"

	"github.com/krishh1at/simple_app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Migration()

	gin.ForceConsoleColor()
	router := gin.Default()
	routes.InitRouter(router)

	s := fmt.Sprintf("/users/%v", 10)
	fmt.Printf("%v : %T", s, s)
	router.Run(":3000")
}
