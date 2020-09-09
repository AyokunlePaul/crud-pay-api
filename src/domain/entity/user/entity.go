package user

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"time"
)

type User struct {
	Id             entity.DatabaseId `json:"id" bson:"_id"`
	FirstName      string            `json:"first_name" bson:"first_name"`
	LastName       string            `json:"last_name" bson:"last_name"`
	Email          string            `json:"email" bson:"email"`
	ProfilePicture string            `json:"profile_picture,omitempty" bson:"profile_picture,omitempty"`
	Password       string            `json:"-" bson:"password"`
	Token          string            `json:"token" bson:"token"`
	RefreshToken   string            `json:"refresh_token" bson:"refresh_token"`
	IsVendor       bool              `json:"is_vendor" bson:"is_vendor"`
	CompanyName    string            `json:"company_name,omitempty" bson:"company_name,omitempty"`
	Phone          string            `json:"phone,omitempty" bson:"phone,omitempty"`
	UserId         string            `json:"user_id" bson:"user_id"`
	CreatedAt      time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at" bson:"updated_at"`
}
