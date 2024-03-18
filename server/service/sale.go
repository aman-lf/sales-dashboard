package service

import (
	"context"
	"strconv"

	"github.com/aman-lf/sales-server/database"
	"github.com/aman-lf/sales-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func GetSalesByProduct(c context.Context, limitStr, offsetStr, searchText string) ([]*model.SalesByProduct, error) {
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 20
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
	}

	pipeline := mongo.Pipeline{
		{{
			Key: "$group",
			Value: bson.D{
				{Key: "_id", Value: "$product_id"},
				{Key: "total_quantity_sold", Value: bson.D{{Key: "$sum", Value: "$quantity"}}},
				{Key: "total_revenue", Value: bson.D{{Key: "$sum", Value: "$total_transaction_amount"}}},
			},
		}},
		{{
			Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: "product"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "product_id"},
				{Key: "as", Value: "product_info"},
			}},
		},
		{{Key: "$unwind", Value: "$product_info"}},
		{{
			Key: "$addFields",
			Value: bson.D{
				{Key: "product_id", Value: "$product_info.product_id"},
				{Key: "product_name", Value: "$product_info.product_name"},
				{Key: "brand_name", Value: "$product_info.brand_name"},
				{Key: "category", Value: "$product_info.category"},
				{Key: "total_profit", Value: bson.D{{Key: "$sum", Value: bson.D{
					{Key: "$multiply", Value: bson.A{"$total_quantity_sold", bson.D{
						{Key: "$subtract", Value: bson.A{"$product_info.selling_price", "$product_info.cost_price"}},
					}},
					}}}}},
			}},
		},
		{{Key: "$sort", Value: bson.D{{Key: "product_id", Value: 1}}}},
		{{Key: "$skip", Value: offset}},
		{{Key: "$limit", Value: limit}},
		{{
			Key: "$project",
			Value: bson.D{
				{Key: "_id", Value: 0}, // Exclude the original _id field if not needed
				{Key: "product_id", Value: 1},
				{Key: "total_quantity_sold", Value: 1},
				{Key: "total_revenue", Value: 1},
				{Key: "total_profit", Value: 1},
				{Key: "product_name", Value: 1},
				{Key: "brand_name", Value: 1},
				{Key: "category", Value: 1},
			}},
		},
	}

	if searchText != "" {
		matchFilter := bson.D{
			{Key: "$or", Value: bson.A{
				bson.D{{Key: "product_name", Value: bson.D{{
					Key:   "$regex",
					Value: primitive.Regex{Pattern: searchText, Options: "i"}}}}},
				bson.D{{Key: "brand_name", Value: bson.D{{
					Key: "$regex", Value: primitive.Regex{Pattern: searchText, Options: "i"}}}}},
				bson.D{{Key: "category", Value: bson.D{{
					Key: "$regex", Value: primitive.Regex{Pattern: searchText, Options: "i"}}}}},
			}},
		}

		pipeline = append(pipeline, bson.D{{Key: "$match", Value: matchFilter}})
	}

	cursor, err := database.FindAggregate(c, model.Sale{}.CollectionName(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	sales := []*model.SalesByProduct{}
	for cursor.Next(c) {
		var sale model.SalesByProduct
		if err := cursor.Decode(&sale); err != nil {
			return nil, err
		}
		sales = append(sales, &sale)
	}

	return sales, nil
}

func GetSalesByBrand(c context.Context, limitStr, offsetStr, searchText string) ([]*model.SalesByBrand, error) {
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 20
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
	}

	pipeline := mongo.Pipeline{
		{{
			Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: "product"},
				{Key: "localField", Value: "product_id"},
				{Key: "foreignField", Value: "product_id"},
				{Key: "as", Value: "product_info"},
			}},
		},
		{{
			Key:   "$unwind",
			Value: "$product_info",
		}},
		{{
			Key: "$group",
			Value: bson.D{
				{Key: "_id", Value: "$product_info.brand_name"},
				{Key: "most_sold_product", Value: bson.D{{Key: "$max", Value: "$product_info.product_name"}}},
				{Key: "total_quantity_sold", Value: bson.D{{Key: "$sum", Value: "$quantity"}}},
				{Key: "total_revenue", Value: bson.D{{Key: "$sum", Value: "$total_transaction_amount"}}},
				{Key: "total_profit", Value: bson.D{{Key: "$sum", Value: bson.D{
					{Key: "$multiply", Value: bson.A{"$quantity", bson.D{
						{Key: "$subtract", Value: bson.A{"$product_info.selling_price", "$product_info.cost_price"}},
					}},
					}}}}},
			},
		}},
		{{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}},
		{{Key: "$skip", Value: offset}},
		{{Key: "$limit", Value: limit}},
		{{
			Key: "$project",
			Value: bson.D{
				{Key: "_id", Value: 0}, // Exclude the original _id field if not needed
				{Key: "brand_name", Value: "$_id"},
				{Key: "most_sold_product", Value: 1},
				{Key: "total_quantity_sold", Value: 1},
				{Key: "total_revenue", Value: 1},
				{Key: "total_profit", Value: 1},
			},
		}},
	}

	if searchText != "" {
		matchFilter := bson.D{
			{Key: "$or", Value: bson.A{
				bson.D{{Key: "product_name", Value: bson.D{{
					Key:   "$regex",
					Value: primitive.Regex{Pattern: searchText, Options: "i"}}}}},
				bson.D{{Key: "brand_name", Value: bson.D{{
					Key: "$regex", Value: primitive.Regex{Pattern: searchText, Options: "i"}}}}},
				bson.D{{Key: "category", Value: bson.D{{
					Key: "$regex", Value: primitive.Regex{Pattern: searchText, Options: "i"}}}}},
			}},
		}

		pipeline = append(pipeline, bson.D{{Key: "$match", Value: matchFilter}})
	}

	cursor, err := database.FindAggregate(c, model.Sale{}.CollectionName(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	sales := []*model.SalesByBrand{}
	for cursor.Next(c) {
		var sale model.SalesByBrand
		if err := cursor.Decode(&sale); err != nil {
			return nil, err
		}
		sales = append(sales, &sale)
	}

	return sales, nil
}
