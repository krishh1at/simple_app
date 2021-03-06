package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"gorm.io/gorm"
)

// User is type of model to capture all database info
// ID            uint64 `gorm:"primaryKey"`
// Name          string `json:"name"`
// GivenName     string `json:"given_name"`
// FamilyName    string `json:"family_name"`
// Profile       string `json:"profile"`
// Picture       string `json:"picture"`
// Email         string `gorm:"unique;not null;size:255"`
// EmailVerified string `json:"email_verified"`
// Gender        string `json:"gender"`
// Admin         bool   `json:"admin"`
// Address       string
// CreatedAt     time.Time
// UpdatedAt     time.Time
// Posts         *[]Post `gorm:"foreignkey:UserID"`
type User struct {
	ID            uint64 `gorm:"primaryKey"`
	Name          string `json:"name"`
	GivenName     string `json:"givenName"`
	FamilyName    string `json:"familyName"`
	Sub           string `json:"sub"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email";gorm:"unique;not null;size:255"`
	EmailVerified bool   `json:"emailVerified"`
	Gender        string `json:"gender"`
	Admin         bool   `json:"admin"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Posts         *[]Post `gorm:"foreignkey:UserID"`
}

// BeforeSave Callback
func (user *User) BeforeSave(db *gorm.DB) error {
	user.Name = user.GivenName + " " + user.FamilyName
	return user.Validate()
}

// BeforeDelete Callback
func (user *User) BeforeDelete(db *gorm.DB) (err error) {
	err = db.Model(&Post{}).Where("user_id = ?", user.ID).Delete(&Post{}).Error

	return err
}

// Validate is for validate user
func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.GivenName, validation.Required, validation.Length(1, 50)),
		validation.Field(&user.FamilyName, validation.Required, validation.Length(1, 50)),
		validation.Field(&user.Email, validation.Required, is.Email, validation.Length(1, 256)),
	)
}

// TableName this method return table name
func (*User) TableName() string {
	return "users"
}

// UpdateAuthDetail to update authentication details
func (user *User) UpdateAuthDetail(u *User, db *gorm.DB) (*User, error) {
	user.Sub = u.Sub

	if u.GivenName != "" {
		user.GivenName = u.GivenName
	}

	if u.FamilyName != "" {
		user.FamilyName = u.FamilyName
	}

	user.EmailVerified = u.EmailVerified
	user.Profile = u.Profile
	user.Picture = u.Picture
	user.Gender = u.Gender
	err := db.Save(user).Error

	return user, err
}
