package routes

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/krishh1at/simple_app/app/controllers"
	"github.com/krishh1at/simple_app/app/controllers/api"
	"github.com/krishh1at/simple_app/app/helpers"
)

// InitRouter is used to setup router
func InitRouter(router *gin.Engine) {
	// assets
	router.Static("/assets/", "./app/assets/")

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

	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "app/views",
		Extension:    ".gohtml",
		Master:       "layouts/application",
		Partials:     []string{"users/form"},
		Funcs:        helpers.Helper,
		DisableCache: true,
	})

	router.GET("/users", controllers.IndexUsers)
	router.GET("/users/:id", controllers.ShowUser)
	router.GET("/user/new", controllers.NewUser)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id/edit", controllers.EditUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id/delete", controllers.DeleteUser)
}
