package db

import (
	"log"

	"github.com/meetsoni1511/go-grpc-auth-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) *Handler {
	log.Println("url", url)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.User{})
	return &Handler{
		db,
	}
}
