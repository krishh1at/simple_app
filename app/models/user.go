package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"gorm.io/gorm"
)

// User is type of model to capture all database info
// Id          uint64    `json:"id"`
// Name        string    `json:"name"`
// Email       string    `json:"email"`
// PhoneNumber string    `json:"phoneNumber"`
// Address     string    `json:"address"`
// Created_at  time.Time `json:"createdAt"`
// Updated_at  time.Time `json:"updatedAt"`
// Posts       has_many   association
type User struct {
	ID          uint64 `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null;size:255"`
	Email       string `gorm:"unique;not null;size:255"`
	PhoneNumber string `gorm:"unique;size:255"`
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Posts       *[]Post `gorm:"foreignkey:UserID"`
}

// const (
// 	AccessKeyID     = "xxxxxxxxxxxxxx"
// 	SecretAccessKey = "xxxxxxxxxxx"
// 	AwsRegion       = "us-east-1"
// )

// TableName this method return table name
func (*User) TableName() string {
	return "users"
}

// BeforeSave Callback
func (user *User) BeforeSave(db *gorm.DB) error {
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
		validation.Field(&user.Name, validation.Required, validation.Length(1, 256)),
		validation.Field(&user.Email, validation.Required, is.Email, validation.Length(1, 256)),
		validation.Field(&user.PhoneNumber, validation.Required, validation.Length(10, 12)),
		validation.Field(&user.Address, validation.Length(1, 256)),
	)
}

// SendSMS is a method to send a message
// func (user *User) SendSMS(msg string) (*sns.PublishOutput, error) {
// 	session, err := session.NewSession(&aws.Config{
// 		Region:      aws.String(AwsRegion),
// 		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	svc := sns.New(session)

// 	params := &sns.PublishInput{
// 		PhoneNumber: aws.String(user.PhoneNumber),
// 		Message:     aws.String(msg),
// 	}

// 	resp, err := svc.Publish(params)

// 	return resp, err
// }
