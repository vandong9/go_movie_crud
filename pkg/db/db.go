
package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return
	}

	return Handler{db}
}