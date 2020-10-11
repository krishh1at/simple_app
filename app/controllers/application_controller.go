package controllers

import (
	"net/http"

	"github.com/krishh1at/simple_app/app/helpers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/krishh1at/simple_app/app/models"
	"github.com/krishh1at/simple_app/config"
)

// CurrentUser is used to find Currentuser
func CurrentUser(c *gin.Context) *models.User {
	var user models.User
	session := sessions.Default(c)
	currentUserID := session.Get("user_id")

	if currentUserID == nil {
		return nil
	} else if err := config.DB.First(&user, currentUserID).Error; err != nil {
		return nil
	}

	return &user
}

// AuthorizedUser to authorize request
func AuthorizedUser(c *gin.Context) {
	var u *models.User
	user := CurrentUser(c)

	if (user == u) || (user == nil) {
		c.Redirect(http.StatusMovedPermanently, helpers.LoginPath())
	}
}
