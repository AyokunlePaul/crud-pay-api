package authentication

import (
	"context"
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/token"
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/user"
	"github.com/AyokunlePaul/crud-pay-api/src/clients/mongo_client"
	"github.com/AyokunlePaul/crud-pay-api/src/clients/sendgrid_client"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/utilities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type userRepository struct {
	tokenRepository token.Repository
}

func NewUserDatabaseRepository(tokenRepository token.Repository) user.Repository {
	return &userRepository{
		tokenRepository: tokenRepository,
	}
}

func (repo *userRepository) CreateUser(user user.User) (*user.User, *response.BaseResponse) {
	user.Id = primitive.NewObjectID()
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient := mongo_client.Get()
	defer func() {
		if closeError := mongoClient.Disconnect(mongoContext); closeError != nil {
			logger.Error("mongo database disconnect error", closeError)
			panic(closeError)
		}
	}()

	user.Password = utilities.GetMD5(user.Password)
	payToken, tokenError := repo.tokenRepository.CreateToken(user.Id.Hex())
	if tokenError != nil {
		return nil, tokenError
	}
	user.Token = payToken.AccessToken
	user.RefreshToken = payToken.RefreshToken

	userCollection := mongoClient.Database("CrudPay").Collection("users")

	_, insertionError := userCollection.InsertOne(mongoContext, user)
	if insertionError != nil {
		return nil, HandleMongoUserExceptions(insertionError)
	}

	return &user, nil
}

func (repo *userRepository) Get(user user.User) (*user.User, *response.BaseResponse) {
	inputPassword := utilities.GetMD5(user.Password)
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient := mongo_client.Get()
	defer func() {
		if closeError := mongoClient.Disconnect(mongoContext); closeError != nil {
			logger.Error("mongo database disconnect error", closeError)
			panic(closeError)
		}
	}()
	userCollection := mongoClient.Database("CrudPay").Collection("users")
	filter := bson.M{"email": user.Email}
	if getUserError := userCollection.FindOne(mongoContext, filter).Decode(&user); getUserError != nil {
		logger.Error("user fetch error", getUserError)
		return nil, HandleMongoUserExceptions(getUserError)
	}
	if passwordError := user.IsValidPassword(inputPassword); passwordError != nil {
		return nil, passwordError
	}

	return &user, nil
}

func (repo *userRepository) Update(user user.User) (*user.User, *response.BaseResponse) {
	return nil, nil
}

func (repo *userRepository) ResetPassword(email string) *response.BaseResponse {
	_ = sendgrid_client.Get()
	return nil
}
