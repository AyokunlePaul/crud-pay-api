package payment

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"time"
)

type Response struct {
	Status  bool
	Message string
	Data    interface{}
}

type Meta struct {
	Total, Skipped, PerPage, Page, PageCount int64
}

type Transaction struct {
	Id        entity.DatabaseId `json:"id" bson:"_id"`
	Name      string            `json:"transaction_name" bson:"name"`
	Amount    float64           `json:"amount" bson:"amount"`
	CreatedAt time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time         `json:"updated_at" bson:"updated_at"`
}
