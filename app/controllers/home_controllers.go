package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeIndex for home routes
func HomeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index", gin.H{
		"title":       "Home",
		"action":      "Index",
		"CurrentPath": c.Request.URL.Path,
	})
}
