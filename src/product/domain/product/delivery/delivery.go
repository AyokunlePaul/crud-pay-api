package delivery

type Group struct {
	Name        string   `json:"group_name" bson:"group_name"`
	Locations   []string `json:"locations" bson:"locations"`
	ShippingFee float64  `json:"shipping_fee" bson:"shipping_fee"`
}
