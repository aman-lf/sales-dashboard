package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sale struct {
	ID                     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TransactionID          string             `json:"transaction_id" bson:"transaction_id"`
	ProductID              string             `json:"product_id" bson:"product_id"`
	Quantity               float64            `json:"quantity" bson:"quantity"`
	TotalTransactionAmount float64            `json:"total_transaction_amount" bson:"total_transaction_amount"`
	TransactionDate        time.Time          `json:"transaction_date" bson:"transaction_date"`
}

type SalesByProduct struct {
	ProductId         string  `json:"id,omitempty" bson:"product_id,omitempty"`
	ProductName       string  `json:"product_name" bson:"product_name"`
	BrandName         string  `json:"brand_name" bson:"brand_name"`
	Category          string  `json:"category" bson:"category"`
	TotalQuantitySold float64 `json:"total_quantity_sold" bson:"total_quantity_sold"`
	TotalRevenue      float64 `json:"total_revenue" bson:"total_revenue"`
	TotalProfit       float64 `json:"total_profit" bson:"total_profit"`
}

type SalesByBrand struct {
	BrandName         string  `json:"brand_name" bson:"brand_name"`
	MostSoldProduct   string  `json:"most_sold_product" bson:"most_sold_product"`
	TotalQuantitySold float64 `json:"total_quantity_sold" bson:"total_quantity_sold"`
	TotalRevenue      float64 `json:"total_revenue" bson:"total_revenue"`
	TotalProfit       float64 `json:"total_profit" bson:"total_profit"`
}

func (s Sale) CollectionName() string {
	return "sale"
}
