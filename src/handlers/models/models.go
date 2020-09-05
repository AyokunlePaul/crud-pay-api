package models

import (
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/user"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"strings"
)

type UserPayload struct {
	Payload map[string]interface{}
}

func (payload *UserPayload) ToDomainUser() user.User {
	domainUser := user.User{}
	if firstName, ok := payload.Payload["first_name"].(string); ok {
		domainUser.FirstName = firstName
	}
	if lastName, ok := payload.Payload["last_name"].(string); ok {
		domainUser.LastName = lastName
	}
	if email, ok := payload.Payload["email"].(string); ok {
		domainUser.Email = email
	}
	if password, ok := payload.Payload["password"].(string); ok {
		domainUser.Password = password
	}
	if isVendor, ok := payload.Payload["is_vendor"].(bool); ok {
		domainUser.IsVendor = isVendor
	}

	return domainUser
}

func (payload *UserPayload) CanBeCreated() *response.BaseResponse {
	if firstName, ok := payload.Payload["first_name"].(string); !ok || strings.TrimSpace(firstName) == "" {
		return response.NewBadRequestError("first name is invalid")
	}
	if lastName, ok := payload.Payload["last_name"].(string); !ok || strings.TrimSpace(lastName) == "" {
		return response.NewBadRequestError("last name is invalid")
	}
	if email, ok := payload.Payload["email"].(string); !ok || strings.TrimSpace(email) == "" {
		return response.NewBadRequestError("email is invalid")
	}
	if password, ok := payload.Payload["password"].(string); !ok || strings.TrimSpace(password) == "" {
		return response.NewBadRequestError("password is invalid")
	}
	return nil
}

func (payload *UserPayload) CanLogin() *response.BaseResponse {
	message := "user_database_repository failed"
	if email, ok := payload.Payload["email"].(string); !ok || strings.TrimSpace(email) == "" {
		return response.NewBadRequestError(message)
	}
	if password, ok := payload.Payload["password"].(string); !ok || strings.TrimSpace(password) == "" {
		return response.NewBadRequestError(message)
	}
	return nil
}
