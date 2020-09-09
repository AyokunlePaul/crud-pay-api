package entity

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatabaseId = primitive.ObjectID
type DefaultID = uuid.UUID
type CrudPayJwtToken = jwt.Token
type SearchResult = elastic.SearchResult

func NewCrudPayId() DatabaseId {
	return primitive.NewObjectID()
}

func StringToCrudPayId(idHex string) (DatabaseId, error) {

	return primitive.ObjectIDFromHex(idHex)
}

func NewDefaultId() DefaultID {
	return uuid.New()
}

func StringToDefaultId(uuidString string) (DefaultID, error) {
	return uuid.Parse(uuidString)
}
