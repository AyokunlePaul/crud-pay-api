package token

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"time"
)

func NewCrudPayToken() *CrudPayToken {
	crudPayToken := new(CrudPayToken)
	crudPayToken.AccessUuid = entity.NewDefaultId().String()
	crudPayToken.RefreshUuid = entity.NewDefaultId().String()
	crudPayToken.AccessTokenExpires = time.Now().Add(3 * 4 * 24 * 7 * time.Hour).Unix()  //3 months
	crudPayToken.RefreshTokenExpires = time.Now().Add(6 * 4 * 24 * 7 * time.Hour).Unix() //6 months

	return crudPayToken
}
