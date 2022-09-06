package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	DB *gorm.DB
)
var Loc, _ = time.LoadLocation("Europe/Minsk")

func NewDB() *gorm.DB {
	dbURL := "postgres://user:pass@database:5432/crud"
	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	return DB
}
func GetDBInstance() *gorm.DB {

	return DB
}
