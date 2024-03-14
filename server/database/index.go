package database

import (
	"context"
	"log"

	"github.com/aman-lf/sales-server/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createUniqueIndex() {
	productIndex := mongo.IndexModel{
		Keys:    bson.M{"product_id": 1},
		Options: options.Index().SetUnique(true),
	}
	saleIndex := mongo.IndexModel{
		Keys:    bson.M{"transaction_id": 1},
		Options: options.Index().SetUnique(true),
	}

	createIndex("product", productIndex)
	createIndex("sale", saleIndex)
}

func createIndex(collectionName string, index mongo.IndexModel) {
	collection := client.Database(config.Cfg.DBName).Collection(collectionName)
	_, err := collection.Indexes().CreateOne(context.Background(), index)
	if err != nil {
		log.Fatal(err)
	}
}
