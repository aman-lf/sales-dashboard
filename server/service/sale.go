package service

import (
	"context"
	"strconv"

	"github.com/aman-lf/sales-server/database"
	"github.com/aman-lf/sales-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetSales(c context.Context, limitStr, offsetStr string) ([]*model.Sale, error) {
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
	sales := []*model.Sale{}

	cursor, err := database.Find(c, model.Sale{}.CollectionName(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var sale model.Sale
		if err := cursor.Decode(&sale); err != nil {
			return nil, err
		}
		sales = append(sales, &sale)
	}

	return sales, nil
}
