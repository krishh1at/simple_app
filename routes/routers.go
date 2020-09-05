package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/krishh1at/simple_app/app/controllers"
)

// InitRouter is used to setup router
func InitRouter(router *gin.Engine) {
	// home routes
	router.GET("/", controllers.HomeIndex)

	// users routes
	api := router.Group("/api/v1")
	api.GET("/users", controllers.IndexUsers)
	api.GET("/users/:id", controllers.ShowUser)
	api.POST("/users", controllers.CreateUser)
	api.PUT("/users/:id", controllers.UpdateUser)
	api.DELETE("/users/:id", controllers.DeleteUser)

	// posts routes
	// api.GET("/users/:userId/posts", controllers.IndexPosts)
	// api.GET("/users/:userId/posts/:id", controllers.ShowPost)
	// api.POST("/users/:userId/posts", controllers.CreatePost)
	// api.PUT("/users/:userId/posts/:id", controllers.UpdatePost)
	// api.DELETE("/users/:userId/posts/:id", controllers.DeletePost)
}
