package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"product-api/config"
	"product-api/model"
	"time"
)

var (
	DB *gorm.DB
)

var Loc, _ = time.LoadLocation("Europe/Minsk")

func NewDB() *gorm.DB {
	conString := config.GetPostgresConnectionString()
	DB, err := gorm.Open(postgres.Open(conString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}

func MigrateProductTable(DB *gorm.DB) {
	c := config.Config{}
	if c.IsMigrate == false {
		err := DB.AutoMigrate(&model.Product{})
		if err != nil {
			panic("migration failed")
		}
	}
}
