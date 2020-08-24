package user

import (
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/utilities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type User struct {
	Id             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName      string             `json:"first_name" bson:"first_name"`
	LastName       string             `json:"last_name" bson:"last_name"`
	Email          string             `json:"email" bson:"email"`
	ProfilePicture string             `json:"profile_picture" bson:"profile_picture"`
	Password       string             `json:"password" bson:"password"`
	Token          string             `json:"token" bson:"token"`
	RefreshToken   string             `json:"refresh_token" bson:"refresh_token"`
}

func (user *User) ValidateUserCreation() *response.BaseResponse {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Password = strings.TrimSpace(user.Password)

	if emailError := user.isValidEmail(); emailError != nil {
		return emailError
	}
	if user.FirstName == "" {
		return response.NewBadRequestError("first name is invalid")
	}
	if user.LastName == "" {
		return response.NewBadRequestError("last name is invalid")
	}
	if user.Password == "" {
		return response.NewBadRequestError("password is invalid")
	}
	return nil
}

func (user *User) ValidateUserLogin() *response.BaseResponse {
	if emailError := user.isValidEmail(); emailError != nil {
		return emailError
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return response.NewBadRequestError("password is invalid")
	}
	return nil
}

func (user *User) IsValidPassword(password string) *response.BaseResponse {
	if user.Password != password {
		return response.NewBadRequestError("authentication error")
	}
	return nil
}

func (user *User) isValidEmail() *response.BaseResponse {
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return response.NewBadRequestError("email is empty")
	}
	if !utilities.IsValidEmail(user.Email) {
		return response.NewBadRequestError("email is invalid")
	}
	return nil
}
