package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishh1at/simple_app/app/models"
	"github.com/krishh1at/simple_app/config"
)

// IndexUsers renders all users
func IndexUsers(c *gin.Context) {
	var users []models.User

	if err := config.DB.Preload("Posts").Order("id desc").Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// ShowUser render corresponding user
func ShowUser(c *gin.Context) {
	user, err := findUser(c)

	if err == nil {
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser to create new user
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser updated corresponding user
func UpdateUser(c *gin.Context) {
	user, err := findUser(c)

	if err != nil {
		return
	}

	c.BindJSON(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser for deleting user of given id
func DeleteUser(c *gin.Context) {
	user, err := findUser(c)

	if err != nil {
		return
	}

	if err = config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "user has been deleted successfully!"})
	}
}

// private function
func findUser(c *gin.Context) (*models.User, error) {
	var (
		user models.User
		err  error
	)

	id := c.Param("id")

	if err = config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	return &user, err
}
