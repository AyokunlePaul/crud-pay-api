package utilities

import (
	"errors"
	"fmt"
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/token"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
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

func CreateToken(userId string) (*token.CrudPayToken, *response.BaseResponse) {
	payToken := token.NewCrudPayToken()
	payToken.AccessUuid = uuid.NewV4().String()
	payToken.RefreshUuid = uuid.NewV4().String()

	accessTokenSecret := os.Getenv(jwtSecret)
	refreshTokenSecret := os.Getenv(jwtRefreshSecret)

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims[accessUuidClaim] = payToken.AccessUuid
	accessTokenClaims[userIdClaim] = userId
	accessTokenClaims[accessTokenClaim] = payToken.AccessTokenExpires

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims[refreshUuidClaim] = payToken.RefreshUuid
	refreshTokenClaims[userIdClaim] = userId
	refreshTokenClaims[refreshTokenClaim] = payToken.RefreshTokenExpires

	var accessTokenError error
	var refreshTokenError error

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	payToken.AccessToken, accessTokenError = accessToken.SignedString([]byte(accessTokenSecret))

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	payToken.RefreshToken, refreshTokenError = refreshToken.SignedString([]byte(refreshTokenSecret))

	if accessTokenError != nil {
		logger.Error("error creating access token", accessTokenError)
		return nil, response.NewInternalServerError(accessTokenError.Error())
	}
	if refreshTokenError != nil {
		logger.Error("error creating refresh token", refreshTokenError)
		return nil, response.NewInternalServerError(refreshTokenError.Error())
	}

	return payToken, nil
}

func VerifyAndExtractToken(accessToken string) (*jwt.Token, *response.BaseResponse) {
	accessTokenSecret := os.Getenv(jwtSecret)
	parsedToken, tokenValidationError := jwt.Parse(accessToken, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			errorMessage := fmt.Sprintf("unexpected signing method: %v", jwtToken.Header["alg"])
			logger.Error("signing method error", errors.New(errorMessage))
			return nil, errors.New(errorMessage)
		}
		return []byte(accessTokenSecret), nil
	})
	if tokenValidationError != nil {
		return nil, response.NewUnAuthorizedError()
	}
	return parsedToken, nil
}

func CheckTokenValidity(accessToken string) *response.BaseResponse {
	parsedToken, parseError := VerifyAndExtractToken(accessToken)
	if parseError != nil {
		return parseError
	}
	if _, ok := parsedToken.Claims.(jwt.Claims); !ok && !parsedToken.Valid {
		return response.NewUnAuthorizedError()
	}
	return nil
}

func GetTokenMetaData(accessToken string) (*string, *response.BaseResponse) {
	parsedToken, parseError := VerifyAndExtractToken(accessToken)
	if parseError != nil {
		return nil, parseError
	}
	tokenClaims, ok := parsedToken.Claims.(jwt.MapClaims)
	if ok && parsedToken.Valid {
		var accessUuid string
		var ok bool
		if accessUuid, ok = tokenClaims[accessUuidClaim].(string); !ok {
			message := fmt.Sprintf("access uuid %v is invalid", tokenClaims[accessUuidClaim])
			logger.Error("invalid access uuid", errors.New(message))
			return nil, response.NewUnAuthorizedError()
		}
		if _, ok = tokenClaims[userIdClaim].(string); !ok {
			message := fmt.Sprintf("user id %v is invalid", tokenClaims[accessUuidClaim])
			logger.Error("invalid user id", errors.New(message))
			return nil, response.NewUnAuthorizedError()
		}
		return &accessUuid, nil
	}

	return nil, response.NewUnAuthorizedError()
}
