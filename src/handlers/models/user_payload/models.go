package user_payload

import (
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/user"
)

type UserPayload struct {
	Payload map[string]interface{}
}

func (payload *UserPayload) ToDomain() user.User {
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
	//Company names are only allowed for vendors
	if domainUser.IsVendor {
		if companyName, ok := payload.Payload["company_name"].(string); ok {
			domainUser.CompanyName = companyName
		}
	}
	if phoneNumber, ok := payload.Payload["phone"].(string); ok {
		domainUser.Phone = phoneNumber
	}

	return domainUser
}
