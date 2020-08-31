package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeIndex for home routes
func HomeIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"home": "index"})
}
