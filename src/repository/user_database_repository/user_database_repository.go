package user_database_repository

import (
	"context"
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/token"
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/user"
	"github.com/AyokunlePaul/crud-pay-api/src/clients/mongo_client"
	"github.com/AyokunlePaul/crud-pay-api/src/clients/sendgrid_client"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/utilities/string_utilities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type userRepository struct {
	tokenRepository token.Repository
}

func New(tokenRepository token.Repository) user.Repository {
	return &userRepository{
		tokenRepository: tokenRepository,
	}
}

func (repo *userRepository) CreateUser(user user.User) (*user.User, *response.BaseResponse) {
	user.Id = primitive.NewObjectID()
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient := mongo_client.Get()

	user.Password = string_utilities.GetMD5(user.Password)
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
	inputPassword := string_utilities.GetMD5(user.Password)
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient := mongo_client.Get()

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

func (repo *userRepository) Update(newUser user.User, token string) (*user.User, *response.BaseResponse) {
	var oldUser user.User
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient := mongo_client.Get()

	userId, userIdError := repo.tokenRepository.Get(token)
	if userIdError != nil {
		return nil, userIdError
	}

	userCollection := mongoClient.Database("CrudPay").Collection("users")
	id, _ := primitive.ObjectIDFromHex(*userId)
	filter := bson.M{
		"_id": id,
	}
	if getUserError := userCollection.FindOne(mongoContext, filter).Decode(&oldUser); getUserError != nil {
		logger.Error("user fetch error", getUserError)
		return nil, HandleMongoUserExceptions(getUserError)
	}

	updateError := oldUser.Update(newUser)
	if updateError != nil {
		return nil, updateError
	}

	updateParameter := bson.D{
		{"$set", bson.D{
			{"first_name", oldUser.FirstName},
			{"last_name", oldUser.LastName},
			{"email", oldUser.Email},
		}},
	}
	if _, updateUserError := userCollection.UpdateOne(mongoContext, filter, updateParameter); updateUserError != nil {
		logger.Error("newUser fetch error", updateUserError)
		return nil, HandleMongoUserExceptions(updateUserError)
	}

	return &oldUser, nil
}

func (repo *userRepository) ResetPassword(email string) *response.BaseResponse {
	_ = sendgrid_client.Get()
	return nil
}

func (repo *userRepository) RefreshToken(refreshToken string) (*user.User, *response.BaseResponse) {
	userResult := new(user.User)
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient := mongo_client.Get()

	newPayToken, tokenError := repo.tokenRepository.RefreshToken(refreshToken)
	if tokenError != nil {
		return nil, tokenError
	}

	userCollection := mongoClient.Database("CrudPay").Collection("users")
	filter := bson.M{
		"refresh_token": refreshToken,
	}

	updateParameter := bson.D{
		{"$set", bson.D{
			{"token", newPayToken.AccessToken},
			{"refresh_token", newPayToken.RefreshToken},
		}},
	}
	if getUserError := userCollection.FindOneAndUpdate(mongoContext, filter, updateParameter).Decode(userResult); getUserError != nil {
		logger.Error("token update error", getUserError)
		return nil, HandleMongoUserExceptions(getUserError)
	}

	return userResult, nil
}
