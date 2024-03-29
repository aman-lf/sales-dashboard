package database

import (
	"context"

	"github.com/aman-lf/sales-server/config"
	"github.com/aman-lf/sales-server/utils/logger"
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
		logger.Fatal(err.Error())
	}

	// Check the connection.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Fatal(err.Error())
	} else {
		logger.Info("Connected to mongoDB!")
	}

	createUniqueIndex()
}

func InsertOne(ctx context.Context, collectionName string, data interface{}) error {
	collection := client.Database(config.Cfg.DBName).Collection(collectionName)

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func Find(ctx context.Context, collectionName string, filter interface{}, findOption *options.FindOptions) (*mongo.Cursor, error) {
	db := client.Database(config.Cfg.DBName)
	collection := db.Collection(collectionName)

	return collection.Find(ctx, filter, findOption)
}

func FindAggregate(ctx context.Context, collectionName string, pipeline interface{}) (*mongo.Cursor, error) {
	db := client.Database(config.Cfg.DBName)
	collection := db.Collection(collectionName)

	return collection.Aggregate(ctx, pipeline)
}
