package user

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/string_utilities"
	"strings"
	"time"
)

func Create() *User {
	newUser := new(User)
	newUser.Id = entity.NewDatabaseId()
	newUser.UserId = entity.NewDefaultId().String()

	currentTime := time.Now()

	newUser.CreatedAt = currentTime
	newUser.UpdatedAt = currentTime

	return newUser
}

func New() *User {
	return new(User)
}

func FromEmailAndPassword(email, password string) *User {
	newUser := new(User)
	newUser.Email = email
	newUser.Password = password

	return newUser
}

func (user *User) CanBeCreated() *response.BaseResponse {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Password = strings.TrimSpace(user.Password)
	user.CompanyName = strings.TrimSpace(user.CompanyName)
	user.Phone = strings.TrimSpace(user.Phone)

	if emailError := user.isValidEmail(); emailError != nil {
		return emailError
	}
	if string_utilities.IsEmpty(user.FirstName) {
		return response.NewBadRequestError("first name is invalid")
	}
	if string_utilities.IsEmpty(user.LastName) {
		return response.NewBadRequestError("last name is invalid")
	}
	if string_utilities.IsEmpty(user.Password) {
		return response.NewBadRequestError("password is invalid")
	}

	if user.IsVendor {
		if string_utilities.IsEmpty(user.CompanyName) {
			return response.NewBadRequestError("company name is invalid")
		}
		if phoneNumberValidationError := user.isValidPhoneNumber(); phoneNumberValidationError != nil {
			return phoneNumberValidationError
		}
	} else {
		if !string_utilities.IsValidPhoneNumber(user.Phone) {
			return response.NewBadRequestError("invalid phone number")
		}
	}

	return nil
}

func (user *User) CanLogin() *response.BaseResponse {
	if emailError := user.isValidEmail(); emailError != nil {
		return emailError
	}
	if passwordError := user.isValidPassword(); passwordError != nil {
		return passwordError
	}
	return nil
}

func (user *User) isValidPassword() *response.BaseResponse {
	user.Password = strings.TrimSpace(user.Password)
	if string_utilities.IsEmpty(user.Password) {
		return response.NewBadRequestError(response.AuthenticationError)
	}
	return nil
}

func (user *User) isValidEmail() *response.BaseResponse {
	user.Email = strings.TrimSpace(user.Email)
	if string_utilities.IsEmpty(user.Email) {
		return response.NewBadRequestError("email is empty")
	}
	if !string_utilities.IsValidEmail(user.Email) {
		return response.NewBadRequestError("email is invalid")
	}
	return nil
}

func (user *User) isValidPhoneNumber() *response.BaseResponse {
	if string_utilities.IsEmpty(user.Phone) {
		return response.NewBadRequestError("phone number is empty")
	}
	if !string_utilities.IsValidPhoneNumber(user.Phone) {
		return response.NewBadRequestError("phone number is invalid")
	}
	return nil
}

func (user *User) CanBeUpdatedWith(newUser User) *response.BaseResponse {
	if !string_utilities.IsEmpty(newUser.FirstName) {
		user.FirstName = newUser.FirstName
	}
	if !string_utilities.IsEmpty(newUser.LastName) {
		user.LastName = newUser.LastName
	}
	if !string_utilities.IsEmpty(newUser.Email) {
		if string_utilities.IsValidEmail(newUser.Email) {
			user.Email = newUser.Email
		} else {
			return response.NewBadRequestError("invalid email")
		}
	}
	if !string_utilities.IsEmpty(newUser.ProfilePicture) {
		user.ProfilePicture = newUser.ProfilePicture
	}
	if !string_utilities.IsEmpty(newUser.Token) {
		user.Token = newUser.Token
	}
	if !string_utilities.IsEmpty(newUser.RefreshToken) {
		user.RefreshToken = newUser.RefreshToken
	}
	return nil
}
