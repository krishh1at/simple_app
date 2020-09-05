package controllers

import (
	"github.com/krishh1at/simple_app/config"
	"github.com/krishh1at/simple_app/models"

	"github.com/gin-gonic/gin"

	"net/http"
)

// UsersIndex renders all users
func UsersIndex(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}

// UserShow render corresponding user
func UserShow(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	config.DB.Find(&user, id)

	c.JSON(http.StatusOK, user)
}

// CreateUser to create new user
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	if err := config.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, user)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser updated corresponding user
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	config.DB.Find(&user, id)
	c.BindJSON(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, user)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser corresponding user
func DeleteUser(c *gin.Context) {
	var user models.User

	id := c.Param("id")

	config.DB.First(&user, id)

	if user.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"User of id " + id: "is not found",
		})
	} else {
		config.DB.Delete(&user)
		c.JSON(http.StatusOK, gin.H{
			"User of id " + id: "is deleted",
		})
	}
}
