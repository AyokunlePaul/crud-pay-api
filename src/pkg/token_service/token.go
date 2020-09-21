package token_service

import (
	"errors"
	"fmt"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
	"github.com/dgrijalva/jwt-go"
	"os"
)

const (
	jwtSecret         = "JWT_SECRET"
	jwtRefreshSecret  = "JWT_SECRET_REFRESH"
	accessUuidClaim   = "access_uuid"
	userIdClaim       = "user_id"
	accessTokenClaim  = "token_expiry"
	refreshUuidClaim  = "refresh_uuid"
	refreshTokenClaim = "access_uuid"
)

type tokenService struct{}

func New() Service {
	return &tokenService{}
}

func (token *tokenService) Create(accessTokenExpires, refreshTokenExpires int64, accessUuid, refreshUuid, userId string) (string, string, *response.BaseResponse) {
	accessTokenSecret := os.Getenv(jwtSecret)
	refreshTokenSecret := os.Getenv(jwtRefreshSecret)

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims[accessUuidClaim] = accessUuid
	accessTokenClaims[userIdClaim] = userId
	accessTokenClaims[accessTokenClaim] = accessTokenExpires

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims[refreshUuidClaim] = refreshUuid
	refreshTokenClaims[userIdClaim] = userId
	refreshTokenClaims[refreshTokenClaim] = refreshTokenExpires

	accessTokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessToken, accessTokenError := accessTokenWithClaims.SignedString([]byte(accessTokenSecret))

	refreshTokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshToken, refreshTokenError := refreshTokenWithClaims.SignedString([]byte(refreshTokenSecret))

	if accessTokenError != nil {
		logger.Error("error creating access token", accessTokenError)
		return "", "", response.NewInternalServerError(accessTokenError.Error())
	}
	if refreshTokenError != nil {
		logger.Error("error creating refresh token", refreshTokenError)
		return "", "", response.NewInternalServerError(refreshTokenError.Error())
	}

	return accessToken, refreshToken, nil
}

func (token *tokenService) VerifyAndExtract(userToken string, isAccessToken bool) (*entity.CrudPayJwtToken, *response.BaseResponse) {
	var tokenSecret string
	if isAccessToken {
		tokenSecret = os.Getenv(jwtSecret)
	} else {
		tokenSecret = os.Getenv(jwtRefreshSecret)
	}

	parsedToken, tokenValidationError := jwt.Parse(userToken, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			errorMessage := fmt.Sprintf("unexpected signing method: %v", jwtToken.Header["alg"])
			logger.Error("signing method error", errors.New(errorMessage))
			return nil, errors.New(errorMessage)
		}
		return []byte(tokenSecret), nil
	})
	if tokenValidationError != nil {
		return nil, response.NewUnAuthorizedError()
	}
	return parsedToken, nil
}

func (token *tokenService) CheckTokenValidity(userToken string, isAccessToken bool) *response.BaseResponse {
	parsedToken, parseError := token.VerifyAndExtract(userToken, isAccessToken)
	if parseError != nil {
		return parseError
	}
	if _, ok := parsedToken.Claims.(jwt.Claims); !ok && !parsedToken.Valid {
		return response.NewUnAuthorizedError()
	}
	return nil
}

func (token *tokenService) GetTokenMetaData(userToken string, isAccessToken bool) (string, *response.BaseResponse) {
	parsedToken, parseError := token.VerifyAndExtract(userToken, isAccessToken)
	if parseError != nil {
		return "", parseError
	}
	tokenClaims, ok := parsedToken.Claims.(jwt.MapClaims)
	if ok && parsedToken.Valid {
		var tokenUuid string
		var ok bool
		if _, ok = tokenClaims[userIdClaim].(string); !ok {
			message := fmt.Sprintf("user id %v is invalid", tokenClaims[accessUuidClaim])
			logger.Error("invalid user id", errors.New(message))
			return "", response.NewUnAuthorizedError()
		}
		if isAccessToken {
			if tokenUuid, ok = tokenClaims[accessUuidClaim].(string); !ok {
				message := fmt.Sprintf("access uuid %v is invalid", tokenClaims[accessUuidClaim])
				logger.Error("invalid access uuid", errors.New(message))
				return "", response.NewUnAuthorizedError()
			}
		} else {
			if tokenUuid, ok = tokenClaims[refreshUuidClaim].(string); !ok {
				message := fmt.Sprintf("refresh uuid %v is invalid", tokenClaims[refreshUuidClaim])
				logger.Error("invalid refresh uuid", errors.New(message))
				return "", response.NewUnAuthorizedError()
			}
		}
		return tokenUuid, nil
	}

	return "", response.NewUnAuthorizedError()
}
