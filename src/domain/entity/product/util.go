package product

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"time"
)

func New() *Product {
	newProduct := new(Product)
	newProduct.Id = entity.NewCrudPayId()
	newProduct.ProductId = entity.NewDefaultId().String()

	currentTime := time.Now()
	newProduct.CreatedAt = currentTime
	newProduct.UpdatedAt = currentTime
	newProduct.DeliveryGroups = []Group{}
	newProduct.Pictures = []string{}

	return newProduct
}
