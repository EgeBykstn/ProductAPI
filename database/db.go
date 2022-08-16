package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"product-api/model"
)

var (
	DB *gorm.DB
)

func NewDB() *gorm.DB {
	dbURL := "postgres://user:pass@localhost:5432/crud"
	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	DB.AutoMigrate(&model.Product{})

	return DB
}
func GetDBInstance() *gorm.DB {

	return DB
}
