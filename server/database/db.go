package database

import (
	"context"
	"fmt"
	"log"

	"github.com/aman-lf/sales-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	ConnectDB()
}

func ConnectDB() {
	var err error
	// Connect to the database.
	clientOptions := options.Client().ApplyURI(config.Cfg.MongoURI)
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to mongoDB!!!")
	}

	createUniqueIndex()
}

func InsertOne(collectionName string, data interface{}) error {
	collection := client.Database(config.Cfg.DBName).Collection(collectionName)

	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}

	return nil
}
