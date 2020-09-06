package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/krishh1at/simple_app/app/controllers"
	"github.com/krishh1at/simple_app/app/controllers/api"
)

// InitRouter is used to setup router
func InitRouter(router *gin.Engine) {
	// images
	router.StaticFile("/images/", "./app/assets/images/")

	// favicon.ico
	router.StaticFile("/favicon.ico", "./app/assets/images/favicon.ico")

	// home routes
	router.GET("/", controllers.HomeIndex)

	// users routes
	users := router.Group("/api")
	users.GET("/users", api.IndexUsers)
	users.GET("/users/:id", api.ShowUser)
	users.POST("/users", api.CreateUser)
	users.PUT("/users/:id", api.UpdateUser)
	users.DELETE("/users/:id", api.DeleteUser)

	// posts routes
	posts := router.Group("/api/user/:user_id/")
	posts.GET("/posts", api.IndexPosts)
	posts.GET("/posts/:id", api.ShowPost)
	posts.POST("/posts", api.CreatePost)
	posts.PUT("/posts/:id", api.UpdatePost)
	posts.DELETE("/posts/:id", api.DeletePost)
}
