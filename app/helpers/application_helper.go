package helpers

import (
	"html/template"
	"strings"

	"github.com/krishh1at/simple_app/config"
)

// Stringify to make string
func Stringify(d interface{}) interface{} {
	var a *string

	if a == d {
		return ""
	}

	return d
}

// ActiveLink to show active link
func ActiveLink(currentPath, path string) string {
	if strings.Contains(currentPath, path) == true {
		return "active"
	}

	return ""
}

// GetLoginURL is used for geting url for oauth2 verification
func GetLoginURL(state string) string {
	return config.Conf.AuthCodeURL(state)
}

// Helper is all user related path
var Helper = template.FuncMap{
	"GetLoginURL":    GetLoginURL,
	"RootPath":       RootPath,
	"UsersPath":      UsersPath,
	"UserPath":       UserPath,
	"NewUserPath":    NewUserPath,
	"EditUserPath":   EditUserPath,
	"UpdateUserPath": UpdateUserPath,
	"DeleteUserPath": DeleteUserPath,
	"ActiveLink":     ActiveLink,
	"string":         Stringify,
}
