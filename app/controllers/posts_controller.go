package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishh1at/simple_app/app/models"
	"github.com/krishh1at/simple_app/config"
)

// IndexPosts for rendering all posts
func IndexPosts(c *gin.Context) {
	user, err := user(c)

	if err != nil {
		return
	}

	var posts []models.Post

	if err := config.DB.Where("user_id = ?", user.ID).Find(&posts).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, posts)
	}
}

// ShowPost is to find post of given id
func ShowPost(c *gin.Context) {
	post, err := findPost(c)

	if err == nil {
		c.JSON(http.StatusOK, post)
	}
}

// CreatePost to create new post
func CreatePost(c *gin.Context) {
	user, err := user(c)

	if err != nil {
		return
	}

	post := models.Post{UserID: &user.ID}
	c.BindJSON(&post)

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, post)
	}
}

// UpdatePost to update post of given id
func UpdatePost(c *gin.Context) {
	post, err := findPost(c)

	if err != nil {
		return
	}

	c.BindJSON(&post)

	if err := config.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, post)
	}
}

// DeletePost to delete post of given id
func DeletePost(c *gin.Context) {
	post, err := findPost(c)

	if err != nil {
		return
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "post has been deleted successfully!"})
	}
}

// private functions
func user(c *gin.Context) (*models.User, error) {
	var (
		user models.User
		err  error
	)

	userID := c.Param("user_id")

	if err = config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user": err.Error()})
	}

	return &user, err
}

func findPost(c *gin.Context) (*models.Post, error) {
	user, err := user(c)
	var post models.Post

	if err == nil {
		post = models.Post{UserID: &user.ID}
		id := c.Param("id")

		if err = config.DB.First(&post, id).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"post": err.Error()})
		}
	}

	return &post, err
}
