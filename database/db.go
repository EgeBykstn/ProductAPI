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
	//migration := config.Migrate{}
	dbURL := "postgres://user:pass@localhost:5432/crud"
	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	//	DB.Migrator().AddColumn(model.Product{}, "CreatedAt")
	//	DB.Migrator().AddColumn(model.Product{}, "UpdatedAt")

	if err != nil {
		log.Fatalln(err)
	}

	if DB.Migrator().HasTable(model.Product{}) == false {
		err := DB.AutoMigrate(&model.Product{})
		if err != nil {
			panic("migration failed")
		}
	}
	return DB
}
func GetDBInstance() *gorm.DB {

	return DB
}
