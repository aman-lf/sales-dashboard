package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID    string             `json:"product_id" bson:"product_id"`
	Name         string             `json:"product_name" bson:"product_name"`
	Brand        string             `json:"brand_name" bson:"brand_name"`
	CostPrice    float64            `json:"cost_price" bson:"cost_price"`
	SellingPrice float64            `json:"selling_price" bson:"selling_price"`
	Category     string             `json:"category" bson:"category"`
	ExpiryDate   time.Time          `json:"expiry_date" bson:"expiry_date"`
}

func (p Product) CollectionName() string {
	return "product"
}
