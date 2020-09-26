package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishh1at/simple_app/app/models"
	"github.com/krishh1at/simple_app/app/paths"
	"github.com/krishh1at/simple_app/config"
)

// IndexUsers renders all users
func IndexUsers(c *gin.Context) {
	var users []models.User

	if err := config.DB.Preload("Posts").Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.HTML(http.StatusOK, "users/index", gin.H{
			"title":  "Users",
			"action": "Index",
			"users":  users,
		})
	}
}

// ShowUser render corresponding user
func ShowUser(c *gin.Context) {
	user, err := findUser(c)

	if err == nil {
		c.JSON(http.StatusOK, user)
	}
}

// NewUser render template for creating new user
func NewUser(c *gin.Context) {
	var user models.User

	c.HTML(http.StatusOK, "users/new", gin.H{
		"title":  "Create New User",
		"action": "Create",
		"user":   user,
	})
}

// CreateUser to create new user
func CreateUser(c *gin.Context) {
	user := userData(c)

	if err := config.DB.Create(&user).Error; err != nil {
		c.HTML(http.StatusOK, "users/new", gin.H{
			"title":  "Create New User",
			"action": "Create",
			"user":   user,
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, paths.UserPath(&user))
	}
}

// EditUser to render update form
func EditUser(c *gin.Context) {
	user, err := findUser(c)

	if err != nil {
		return
	} else {
		c.HTML(http.StatusOK, "users/edit", gin.H{
			"title":  "Update User",
			"action": "Update",
			"user":   user,
		})
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
		c.HTML(http.StatusOK, "users/edit", gin.H{
			"title":  "Update User",
			"action": "Update",
			"user":   user,
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, paths.UserPath(user))
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

func userData(c *gin.Context) models.User {
	name := c.PostForm("user[Name]")
	email := c.PostForm("user[Email]")
	phone := c.PostForm("user[PhoneNumber]")

	user := models.User{
		Name:        &name,
		Email:       &email,
		PhoneNumber: &phone,
		Address:     c.PostForm("user[Address]"),
	}

	return user
}
