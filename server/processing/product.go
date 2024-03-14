package processing

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/aman-lf/sales-server/database"
	"github.com/aman-lf/sales-server/model"
	"github.com/aman-lf/sales-server/utils"
)

func ProcessProducts(path string) {
	sleepTime := 90

	for i := 1; ; i++ {
		filename := fmt.Sprintf("products_%d.csv", i)
		fmt.Println("Reading product file:", filename)

		filePath := filepath.Join(path, filename)
		err := processProductData(filePath)

		if err != nil {
			i--
			time.Sleep(time.Duration(sleepTime) * time.Second)
		}
	}
}

func processProductData(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		fmt.Printf("Error reading CSV header: %v\n", err)
		return err
	}

	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file
		}

		var product model.Product
		for i, field := range record {
			switch header[i] {
			case "product_id":
				product.ProductID = field
			case "product_name":
				product.Name = field
			case "brand_name":
				product.Brand = field
			case "cost_price":
				product.CostPrice = utils.ParseFloat(field)
			case "selling_price":
				product.SellingPrice = utils.ParseFloat(field)
			case "category":
				product.Category = field
			case "expiry_date":
				product.ExpiryDate, _ = time.Parse("2006-01-02", field)
			}
		}

		database.InsertOne(product.CollectionName(), product)
	}

	return nil
}
