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

func (s Sale) CollectionName() string {
	return "sale"
}
