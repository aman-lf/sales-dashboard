package service

import (
	"github.com/aman-lf/sales-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCountPipeline(groupBy string, filter *model.PipelineParams) mongo.Pipeline {
	var groupValue string
	switch groupBy {
	case PRODUCT:
		groupValue = "$product_info.product_id"
	case BRAND:
		groupValue = "$product_info.brand_name"
	}

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
				{Key: "_id", Value: groupValue},
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

	if filter.SearchText != "" {
		orCase := bson.A{
			bson.D{{Key: "product_info.product_name", Value: bson.D{{
				Key:   "$regex",
				Value: primitive.Regex{Pattern: filter.SearchText, Options: "i"}}}}},
			bson.D{{Key: "product_info.brand_name", Value: bson.D{{
				Key: "$regex", Value: primitive.Regex{Pattern: filter.SearchText, Options: "i"}}}}},
		}
		if groupBy == PRODUCT {
			orCase = append(orCase, bson.D{{Key: "product_info.category", Value: bson.D{{
				Key: "$regex", Value: primitive.Regex{Pattern: filter.SearchText, Options: "i"}}}}})
		}

		countMatchFilter := bson.D{{
			Key: "$match",
			Value: bson.D{
				{Key: "$or", Value: orCase},
			},
		}}
		insertIndex := 2
		countPipeline = append(countPipeline[:insertIndex], append([]bson.D{countMatchFilter}, countPipeline[insertIndex:]...)...)
	}

	return countPipeline
}

func getSalesByProductPipeline(filter *model.PipelineParams) mongo.Pipeline {
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
		{{
			Key: "$project",
			Value: bson.D{
				{Key: "_id", Value: 0}, // Exclude the original _id field
				{Key: "product_id", Value: 1},
				{Key: "total_quantity_sold", Value: 1},
				{Key: "total_revenue", Value: 1},
				{Key: "total_profit", Value: 1},
				{Key: "product_name", Value: 1},
				{Key: "brand_name", Value: 1},
				{Key: "category", Value: 1},
			}},
		},
		{{Key: "$sort", Value: bson.D{{Key: filter.SortBy, Value: filter.SortOrder}}}},
		{{Key: "$skip", Value: filter.Offset}},
		{{Key: "$limit", Value: filter.Limit}},
	}

	if filter.SearchText != "" {
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
		insertIndex := 5
		pipeline = append(pipeline[:insertIndex], append([]bson.D{matchFilter}, pipeline[insertIndex:]...)...)
	}

	return pipeline
}

func getSalesByBrandPipeline(filter *model.PipelineParams) mongo.Pipeline {
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
					}}},
				}}}},
			},
		}},
		{{
			Key: "$project",
			Value: bson.D{
				{Key: "_id", Value: 0}, // Exclude the original _id field
				{Key: "brand_name", Value: "$_id"},
				{Key: "most_sold_product", Value: 1},
				{Key: "total_quantity_sold", Value: 1},
				{Key: "total_revenue", Value: 1},
				{Key: "total_profit", Value: 1},
			},
		}},
		{{Key: "$sort", Value: bson.D{{Key: filter.SortBy, Value: filter.SortOrder}}}},
		{{Key: "$skip", Value: filter.Offset}},
		{{Key: "$limit", Value: filter.Limit}},
	}

	if filter.SearchText != "" {
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
		insertIndex := 4
		pipeline = append(pipeline[:insertIndex], append([]bson.D{matchFilter}, pipeline[insertIndex:]...)...)
	}

	return pipeline
}

func getProductProfitPipeline(pType string) mongo.Pipeline {
	var order int
	switch pType {
	case MOST:
		order = -1
	case LEAST:
		order = 1
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
				{Key: "_id", Value: "$product_info.product_name"},
				{Key: "total_profit", Value: bson.D{{Key: "$sum", Value: bson.D{
					{Key: "$multiply", Value: bson.A{"$quantity", bson.D{
						{Key: "$subtract", Value: bson.A{"$product_info.selling_price", "$product_info.cost_price"}},
					}}},
				}}}},
			},
		}},
		{{
			Key: "$project",
			Value: bson.D{
				{Key: "_id", Value: 0}, // Exclude the original _id field
				{Key: "product_name", Value: "$_id"},
				{Key: "total_profit", Value: 1},
			},
		}},
		{{Key: "$sort", Value: bson.D{{Key: "total_profit", Value: order}}}},
		{{Key: "$limit", Value: 1}},
	}

	return pipeline
}

func salesDatePipeline(dType string) mongo.Pipeline {
	var order int
	switch dType {
	case MOST:
		order = -1
	case LEAST:
		order = 1
	}

	pipeline := mongo.Pipeline{
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
		{{
			Key: "$group",
			Value: bson.D{
				{Key: "_id", Value: "$transaction_date"},
				{Key: "total_sales", Value: bson.D{{Key: "$sum", Value: "$quantity"}}},
			},
		}},
		{{
			Key: "$project",
			Value: bson.D{
				{Key: "_id", Value: 0}, // Exclude the original _id field
				{Key: "transaction_date", Value: "$_id"},
				{Key: "total_sales", Value: 1},
			},
		}},
		{{
			Key:   "$sort",
			Value: bson.D{{Key: "total_sales", Value: order}},
		}},
		{{
			Key:   "$limit",
			Value: 1,
		}},
	}

	return pipeline
}
