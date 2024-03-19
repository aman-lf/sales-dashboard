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

func GetSalesByProduct(c context.Context, filter model.PipelineParams) ([]*model.SalesByProduct, error) {
	countPipeline := mongo.Pipeline{
		{{
			Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: "product"},
				{Key: "localField", Value: "product_id"},
				{Key: "foreignField", Value: "product_id"},
				{Key: "as", Value: "product_info"},
			},
		}},
		{{
			Key:   "$unwind",
			Value: "$product_info",
		}},
		// Add match filter here
		{{
			Key: "$group",
			Value: bson.D{
				{Key: "_id", Value: "$product_info.product_id"},
			},
		}},
		{{
			Key: "$group",
			Value: bson.D{
				{Key: "_id", Value: nil},
				{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
			},
		}},
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
		{{Key: "$sort", Value: bson.D{{Key: filter.SortBy, Value: filter.SortOrder}}}},
		{{Key: "$skip", Value: filter.Offset}},
		{{Key: "$limit", Value: filter.Limit}},
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

	if filter.SearchText != "" {
		countMatchFilter := bson.D{{
			Key: "$match",
			Value: bson.D{
				{Key: "$or", Value: bson.A{
					bson.D{{Key: "product_info.product_name", Value: bson.D{{
						Key:   "$regex",
						Value: primitive.Regex{Pattern: filter.SearchText, Options: "i"}}}}},
					bson.D{{Key: "product_info.brand_name", Value: bson.D{{
						Key: "$regex", Value: primitive.Regex{Pattern: filter.SearchText, Options: "i"}}}}},
					bson.D{{Key: "product_info.category", Value: bson.D{{
						Key: "$regex", Value: primitive.Regex{Pattern: filter.SearchText, Options: "i"}}}}},
				}},
			},
		}}
		insertIndex := 2
		countPipeline = append(countPipeline[:insertIndex], append([]bson.D{countMatchFilter}, countPipeline[insertIndex:]...)...)

		matchFilter := bson.D{{
			Key: "$match",
			Value: bson.D{
				{Key: "$or", Value: bson.A{
					bson.D{{Key: "product_name", Value: bson.D{{
						Key:   "$regex",
						Value: primitive.Regex{Pattern: filter.SearchText, Options: "i"}}}}},
					bson.D{{Key: "brand_name", Value: bson.D{{
						Key: "$regex", Value: primitive.Regex{Pattern: filter.SearchText, Options: "i"}}}}},
					bson.D{{Key: "category", Value: bson.D{{
						Key: "$regex", Value: primitive.Regex{Pattern: filter.SearchText, Options: "i"}}}}},
				}},
			},
		}}
		pipeline = append(pipeline, matchFilter)
	}

	countCursor, err := database.FindAggregate(c, model.Sale{}.CollectionName(), countPipeline)
	if err != nil {
		return nil, err
	}
	defer countCursor.Close(c)

	var countResult struct {
		Count int `bson:"count"`
	}
	for countCursor.Next(c) {
		if err := countCursor.Decode(&countResult); err != nil {
			return nil, err
		}
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

func GetSalesByBrand(c context.Context, filter model.PipelineParams) ([]*model.SalesByBrand, error) {
	countPipeline := mongo.Pipeline{
		{{
			Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: "product"},
				{Key: "localField", Value: "product_id"},
				{Key: "foreignField", Value: "product_id"},
				{Key: "as", Value: "product_info"},
			},
		}},
		{{
			Key:   "$unwind",
			Value: "$product_info",
		}},
		// Add match filter here
		{{
			Key: "$group",
			Value: bson.D{
				{Key: "_id", Value: "$product_info.brand_name"},
			},
		}},
		{{
			Key: "$group",
			Value: bson.D{
				{Key: "_id", Value: nil},
				{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
			},
		}},
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
		{{Key: "$sort", Value: bson.D{{Key: filter.SortBy, Value: filter.SortOrder}}}},
		{{Key: "$skip", Value: filter.Offset}},
		{{Key: "$limit", Value: filter.Limit}},
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

	if filter.SearchText != "" {
		countMatchFilter := bson.D{{
			Key: "$match",
			Value: bson.D{
				{Key: "$or", Value: bson.A{
					bson.D{{Key: "product_info.product_name", Value: bson.D{{
						Key:   "$regex",
						Value: primitive.Regex{Pattern: filter.SearchText, Options: "i"}}}}},
					bson.D{{Key: "product_info.brand_name", Value: bson.D{{
						Key: "$regex", Value: primitive.Regex{Pattern: filter.SearchText, Options: "i"}}}}},
				}},
			},
		}}
		insertIndex := 2
		countPipeline = append(countPipeline[:insertIndex], append([]bson.D{countMatchFilter}, countPipeline[insertIndex:]...)...)

		matchFilter := bson.D{{
			Key: "$match",
			Value: bson.D{
				{Key: "$or", Value: bson.A{
					bson.D{{Key: "most_sold_product", Value: bson.D{{
						Key:   "$regex",
						Value: primitive.Regex{Pattern: filter.SearchText, Options: "i"}}}}},
					bson.D{{Key: "brand_name", Value: bson.D{{
						Key: "$regex", Value: primitive.Regex{Pattern: filter.SearchText, Options: "i"}}}}},
				}},
			},
		}}
		pipeline = append(pipeline, matchFilter)
	}

	countCursor, err := database.FindAggregate(c, model.Sale{}.CollectionName(), countPipeline)
	if err != nil {
		return nil, err
	}
	defer countCursor.Close(c)

	var countResult struct {
		Count int `bson:"count"`
	}
	for countCursor.Next(c) {
		if err := countCursor.Decode(&countResult); err != nil {
			return nil, err
		}
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
