package service

import (
	"context"
	"strconv"
	"time"

	"github.com/aman-lf/sales-server/database"
	"github.com/aman-lf/sales-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	PRODUCT              = "product"
	BRAND                = "brand"
	PRODUCT_DEFAULT_SORT = "product_id"
	BRAND_DEFAULT_SORT   = "brand_name"
	MOST                 = "most"
	LEAST                = "least"
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

func GetSalesByProduct(c context.Context, filter *model.PipelineParams) ([]*model.SalesByProduct, error) {
	countPipeline := getCountPipeline(PRODUCT, filter)
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

	filter.CalculateTotalPageCount(countResult.Count)
	filter.VerifyPage()
	filter.CalculateOffset()

	pipeline := getSalesByProductPipeline(filter)
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

func GetSalesByBrand(c context.Context, filter *model.PipelineParams) ([]*model.SalesByBrand, error) {
	countPipeline := getCountPipeline(BRAND, filter)
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

	filter.CalculateTotalPageCount(countResult.Count)
	filter.VerifyPage()
	filter.CalculateOffset()

	pipeline := getSalesByBrandPipeline(filter)
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

func GetDashboardData(c context.Context) (*model.SaleDetails, error) {
	type profitData struct {
		ProductName string  `bson:"product_name"`
		TotalProfit float64 `bson:"total_profit"`
	}

	mostProfitPipeline := getProductProfitPipeline(MOST)
	cursor, err := database.FindAggregate(c, model.Sale{}.CollectionName(), mostProfitPipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var mostProfit profitData
	for cursor.Next(c) {
		if err := cursor.Decode(&mostProfit); err != nil {
			return nil, err
		}
	}

	leastProfitPipeline := getProductProfitPipeline(LEAST)
	cursor, err = database.FindAggregate(c, model.Sale{}.CollectionName(), leastProfitPipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var leastProfit profitData
	for cursor.Next(c) {
		if err := cursor.Decode(&leastProfit); err != nil {
			return nil, err
		}
	}

	type saleData struct {
		TransactionDate time.Time `bson:"transaction_date"`
		TotalSales      int       `bson:"total_sales"`
	}

	mostSalePipeline := salesDatePipeline(MOST)
	cursor, err = database.FindAggregate(c, model.Sale{}.CollectionName(), mostSalePipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var mostSale saleData
	for cursor.Next(c) {
		if err := cursor.Decode(&mostSale); err != nil {
			return nil, err
		}
	}

	leastSalePipeline := salesDatePipeline(LEAST)
	cursor, err = database.FindAggregate(c, model.Sale{}.CollectionName(), leastSalePipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var leastSale saleData
	for cursor.Next(c) {
		if err := cursor.Decode(&leastSale); err != nil {
			return nil, err
		}
	}

	mostSaletransactionDate := mostSale.TransactionDate.Format("2006-01-02")
	leastSaletransactionDate := leastSale.TransactionDate.Format("2006-01-02")

	return &model.SaleDetails{
		MostProfitableProduct:  mostProfit.ProductName,
		LeastProfitableProduct: leastProfit.ProductName,
		HighestSalesDate:       mostSaletransactionDate,
		LeastSalesDate:         leastSaletransactionDate,
	}, nil
}
