package app

import (
	"log"
	"os"
	"path/filepath"

	"github.com/aman-lf/sales-server/config"
	"github.com/aman-lf/sales-server/processing"
)

func ProcessFiles() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	productPath := filepath.Join(workingDir, config.Cfg.FilePath, "products")
	salePath := filepath.Join(workingDir, config.Cfg.FilePath, "sales")
	go processing.ProcessProducts(productPath)
	go processing.ProcessSales(salePath)
}
