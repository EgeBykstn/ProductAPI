package database

import (
	"gorm.io/gorm"
	"log"
	"product-api/config"
)

var (
	DB *gorm.DB
)

func NewDB() *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	DB, err = gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
