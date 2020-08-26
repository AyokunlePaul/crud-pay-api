package user

import (
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/utilities/string_utilities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type User struct {
	Id             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName      string             `json:"first_name" bson:"first_name"`
	LastName       string             `json:"last_name" bson:"last_name"`
	Email          string             `json:"email" bson:"email"`
	ProfilePicture string             `json:"profile_picture,omitempty" bson:"profile_picture"`
	Password       string             `json:"-" bson:"password"`
	Token          string             `json:"token" bson:"token"`
	RefreshToken   string             `json:"refresh_token" bson:"refresh_token"`
	IsVendor       bool               `json:"is_vendor" bson:"is_vendor"`
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
	if !string_utilities.IsValidEmail(user.Email) {
		return response.NewBadRequestError("email is invalid")
	}
	return nil
}

func (user *User) Update(newUser User) *response.BaseResponse {
	if !string_utilities.IsEmpty(strings.TrimSpace(newUser.FirstName)) {
		user.FirstName = newUser.FirstName
	}
	if !string_utilities.IsEmpty(strings.TrimSpace(newUser.LastName)) {
		user.LastName = newUser.LastName
	}
	if !string_utilities.IsEmpty(strings.TrimSpace(newUser.Email)) {
		if string_utilities.IsValidEmail(newUser.Email) {
			user.Email = newUser.Email
		} else {
			return response.NewBadRequestError("invalid email")
		}
	}
	if !string_utilities.IsEmpty(strings.TrimSpace(newUser.ProfilePicture)) {
		user.ProfilePicture = newUser.ProfilePicture
	}
	return nil
}
