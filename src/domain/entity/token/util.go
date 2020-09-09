package token

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"time"
)

func NewCrudPayToken() *CrudPayToken {
	crudPayToken := new(CrudPayToken)
	crudPayToken.AccessUuid = entity.NewDefaultId().String()
	crudPayToken.RefreshUuid = entity.NewDefaultId().String()
	crudPayToken.AccessTokenExpires = time.Now().Add(24 * 7 * time.Hour).Unix()      //7 days
	crudPayToken.RefreshTokenExpires = time.Now().Add(4 * 24 * 7 * time.Hour).Unix() //1 month

	return crudPayToken
}
