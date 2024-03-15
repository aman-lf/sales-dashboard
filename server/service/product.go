package service

import (
	"context"
	"strconv"

	"github.com/aman-lf/sales-server/database"
	"github.com/aman-lf/sales-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProducts(c context.Context, limitStr, offsetStr string) ([]*model.Product, error) {
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 20
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
	}

	findOptions := options.Find().SetLimit(limit).SetSkip(offset).SetSort(bson.D{{Key: "_id", Value: 1}})

	filter := bson.M{}
	products := []*model.Product{}

	cursor, err := database.Find(c, model.Product{}.CollectionName(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var product model.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}
