package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/krishh1at/simple_app/app/helpers"
	"github.com/krishh1at/simple_app/app/models"
	"github.com/krishh1at/simple_app/config"
	"golang.org/x/oauth2"
)

var session sessions.Session

func randomToken() string {
	b := make([]byte, 32)
	rand.Read(b)

	return base64.StdEncoding.EncodeToString(b)
}

// LoginHandler to render login template
func LoginHandler(c *gin.Context) {
	config.State = randomToken()
	session := sessions.Default(c)
	session.Set("state", config.State)
	session.Save()

	c.HTML(http.StatusOK, "sessions/new", gin.H{
		"title":       "Users",
		"action":      "Login",
		"CurrentPath": c.Request.URL.Path,
		"state":       config.State,
	})
}

// AuthHandler to authenticate
func AuthHandler(c *gin.Context) {
	session := sessions.Default(c)
	retrievedState := session.Get("state")

	if retrievedState != c.Query("state") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := config.Conf.Exchange(oauth2.NoContext, c.Query("code"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := config.Conf.Client(oauth2.NoContext, token)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	defer email.Body.Close()

	data, _ := ioutil.ReadAll(email.Body)
	user, err := authUser(data)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID := user.ID
	session.Set("user_id", userID)
	session.Save()

	c.Redirect(http.StatusMovedPermanently, helpers.UsersPath())
}

// LogoutHandler to render login template
func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user_id", nil)
	session.Delete("user_id")
	session.Save()

	c.Redirect(http.StatusMovedPermanently, helpers.LoginPath())
}

func authUser(data []byte) (*models.User, error) {
	var dbUser, authUser models.User

	err := json.Unmarshal(data, &authUser)

	if err != nil {
		return nil, err
	}

	authEmail := authUser.Email

	if err := config.DB.Where("email = ?", authEmail).First(&dbUser).Error; err != nil {
		return nil, err
	}

	return dbUser.UpdateAuthDetail(&authUser, config.DB)
}
