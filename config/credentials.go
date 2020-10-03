package config

import (
	"io/ioutil"
	"log"

	"github.com/gin-contrib/sessions/cookie"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gopkg.in/yaml.v2"
)

type Credential struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

var cred Credential
var Conf *oauth2.Config
var State string
var Store = cookie.NewStore([]byte("secret"))

func init() {
	file, err := ioutil.ReadFile("./config/credentials.yml")

	if err != nil {
		log.Panic(err)
	}

	yaml.Unmarshal(file, &cred)

	Conf = &oauth2.Config{
		ClientID:     cred.ClientID,
		ClientSecret: cred.ClientSecret,
		RedirectURL:  "http://localhost:8080/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}
