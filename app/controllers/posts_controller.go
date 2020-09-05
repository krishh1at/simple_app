package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishh1at/simple_app/app/models"
	"github.com/krishh1at/simple_app/config"
)

// IndexPosts for rendering all posts
func IndexPosts(c *gin.Context) {
	var posts []models.Post

	if err := config.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, posts)
	}
}

// ShowPost is to find post of given id
func ShowPost(c *gin.Context) {
	var post models.Post

	id := c.Param("id")

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, post)
	}
}

// CreatePost to create new post
func CreatePost(c *gin.Context) {
	var post models.Post

	c.BindJSON(&post)

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, post)
	}
}

// UpdatePost to update post of given id
func UpdatePost(c *gin.Context) {
	var post models.Post

	id := c.Param("id")

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
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
	var post models.Post
	id := c.Param("id")

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "post has been deleted successfully!"})
	}
}
