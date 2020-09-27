package helpers

import (
	"html/template"
	"strings"
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

// Helper is all user related path
var Helper = template.FuncMap{
	"RootPath":       RootPath,
	"UsersPath":      UsersPath,
	"UserPath":       UserPath,
	"NewUserPath":    NewUserPath,
	"EditUserPath":   EditUserPath,
	"DeleteUserPath": DeleteUserPath,
	"ActiveLink":     ActiveLink,
	"string":         Stringify,
}
