package db

import (
	"github.com/meetsoni1511/go-grpc-product-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Product{}, &models.StockDecreaseLog{})
	db.AutoMigrate(&models.StockDecreaseLog{})
	return Handler{db}
}
