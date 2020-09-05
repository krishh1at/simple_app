package routers

import (
	"github.com/krishh1at/simple_app/controllers"

	"github.com/gin-gonic/gin"
)

// InitRouter is used to setup router
func InitRouter(router *gin.Engine) {
	router.GET("/", controllers.HomeIndex)
	router.GET("/users", controllers.UsersIndex)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.UserShow)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
}
