package token

import "time"

type CrudPayToken struct {
	AccessToken         string
	RefreshToken        string
	AccessUuid          string
	RefreshUuid         string
	AccessTokenExpires  int64
	RefreshTokenExpires int64
}

func NewCrudPayToken() *CrudPayToken {
	return &CrudPayToken{
		AccessTokenExpires:  time.Now().Add(5 * time.Minute).Unix(),
		RefreshTokenExpires: time.Now().Add(24 * time.Hour).Unix(),
	}
}

func (token *CrudPayToken) TokenIsExpired() bool {
	return time.Now().After(time.Unix(token.AccessTokenExpires, 0))
}

func (token *CrudPayToken) RefreshTokenIsExpired() bool {
	return time.Now().After(time.Unix(token.RefreshTokenExpires, 0))
}
