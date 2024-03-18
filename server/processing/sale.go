package processing

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/aman-lf/sales-server/database"
	"github.com/aman-lf/sales-server/model"
	"github.com/aman-lf/sales-server/service"
	"github.com/aman-lf/sales-server/utils"
)

func ProcessSales(path string) {
	sleepTime := 60

	for i := 1; ; i++ {
		filename := fmt.Sprintf("sales_%d.csv", i)

		filePath := filepath.Join(path, filename)
		err := processSalesData(filePath)
		if err != nil {
			i--
			time.Sleep(time.Duration(sleepTime) * time.Second)
		} else {
			fmt.Println("Processed sales file:", filename)
			service.TriggerSSEMessages("New data is available! Please refresh the page.")
		}
	}
}

func processSalesData(filepath string) error {
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

		var sale model.Sale
		for i, field := range record {
			switch header[i] {
			case "transaction_id":
				sale.TransactionID = field
			case "product_id":
				sale.ProductID = field
			case "quantity":
				sale.Quantity = utils.ParseFloat(field)
			case "total_transacion_amount":
				sale.TotalTransactionAmount = utils.ParseFloat(field)
			case "transaction_date":
				sale.TransactionDate, _ = time.Parse("2006-01-02", field)
			}
		}

		database.InsertOne(context.Background(), sale.CollectionName(), sale)
	}

	return nil
}
