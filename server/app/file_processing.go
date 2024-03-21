package app

import (
	"os"
	"path/filepath"

	"github.com/aman-lf/sales-server/config"
	"github.com/aman-lf/sales-server/processing"
	"github.com/aman-lf/sales-server/utils/logger"
)

func ProcessFiles() {
	workingDir, err := os.Getwd()
	if err != nil {
		logger.Fatal(err.Error())
	}

	productPath := filepath.Join(workingDir, config.Cfg.FilePath, "products")
	salePath := filepath.Join(workingDir, config.Cfg.FilePath, "sales")
	go processing.ProcessProducts(productPath)
	go processing.ProcessSales(salePath)
}
