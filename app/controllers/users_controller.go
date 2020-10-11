package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishh1at/simple_app/app/helpers"
	"github.com/krishh1at/simple_app/app/models"
	"github.com/krishh1at/simple_app/config"
)

// IndexUsers renders all users
func IndexUsers(c *gin.Context) {
	var users []models.User

	if err := config.DB.Preload("Posts").Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.HTML(http.StatusOK, "users/index", gin.H{
			"title":       "Users",
			"action":      "Index",
			"users":       users,
			"CurrentPath": c.Request.URL.Path,
		})
	}
}

// ShowUser render corresponding user
func ShowUser(c *gin.Context) {
	user, err := findUser(c)

	if err == nil {
		c.HTML(http.StatusOK, "users/show", gin.H{
			"title":       "User",
			"action":      "Show",
			"user":        &user,
			"CurrentPath": c.Request.URL.Path,
		})
	}
}

// NewUser render template for creating new user
func NewUser(c *gin.Context) {
	AuthorizedUser(c)

	var user models.User

	c.HTML(http.StatusOK, "users/new", gin.H{
		"title":       "Create New User",
		"action":      "Create",
		"user":        user,
		"CurrentPath": c.Request.URL.Path,
	})
}

// CreateUser to create new user
func CreateUser(c *gin.Context) {
	AuthorizedUser(c)

	var u models.User
	user, err := BindFormData(c, u)

	if err != nil {
		return
	}

	if err := config.DB.Create(user).Error; err != nil {
		c.HTML(http.StatusOK, "users/new", gin.H{
			"title":       "Create New User",
			"action":      "Create",
			"user":        user,
			"errors":      user.Validate(),
			"CurrentPath": c.Request.URL.Path,
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, helpers.UserPath(user))
	}
}

// EditUser to render update form
func EditUser(c *gin.Context) {
	user := AuthorizedUser(c)

	c.HTML(http.StatusOK, "users/edit", gin.H{
		"title":       "Update User",
		"action":      "Update",
		"user":        &user,
		"CurrentPath": c.Request.URL.Path,
	})
}

// UpdateUser updated corresponding user
func UpdateUser(c *gin.Context) {
	u := AuthorizedUser(c)
	user, err := BindFormData(c, *u)

	if err != nil {
		return
	}

	if err := config.DB.Save(user).Error; err != nil {
		c.HTML(http.StatusOK, "users/edit", gin.H{
			"title":       "Update User",
			"action":      "Update",
			"user":        user,
			"errors":      user.Validate(),
			"CurrentPath": c.Request.URL.Path,
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, helpers.UserPath(user))
	}
}

// DeleteUser for deleting user of given id
func DeleteUser(c *gin.Context) {
	authUser := AuthorizedUser(c)
	user, err := findUser(c)

	if err != nil {
		return
	}

	if authUser.ID == user.ID {
		if err := config.DB.Delete(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, "Error")
		}
	}

	c.Redirect(http.StatusMovedPermanently, helpers.UsersPath())
}

// private function
func findUser(c *gin.Context) (models.User, error) {
	var (
		user models.User
		err  error
	)

	id := c.Param("id")

	if err = config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	return user, err
}

// BindFormData to bind form data
func BindFormData(c *gin.Context, user models.User) (*models.User, error) {
	userData, _ := c.GetPostFormMap("user")
	jsonbody, _ := json.Marshal(userData)
	_ = json.Unmarshal(jsonbody, &user)

	return &user, nil
}
