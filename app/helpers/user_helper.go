package helpers

import (
	"fmt"

	"github.com/krishh1at/simple_app/app/models"
)

// UserPath is path to their show page
func UserPath(user *models.User) string {
	return fmt.Sprintf("/users/%v", user.ID)
}

// UsersPath is users index path
func UsersPath() string {
	return "/users"
}

// NewUserPath is path to create new user
func NewUserPath() string {
	return "/user/new"
}

// EditUserPath to render edit user page
func EditUserPath(user *models.User) string {
	return fmt.Sprintf("/users/%v/edit", user.ID)
}

// DeleteUserPath to delete user of given id
func DeleteUserPath(user *models.User) string {
	return fmt.Sprintf("/users/%v/delete", user.ID)
}
