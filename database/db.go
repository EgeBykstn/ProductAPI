package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"product-api/config"
	"product-api/model"
	"time"
)

var (
	DB *gorm.DB
)

var Loc, _ = time.LoadLocation("Europe/Minsk")

type GormDB interface {
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Model(value interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
}

func NewDB() *gorm.DB {
	conString := config.GetPostgresConnectionString()
	DB, err := gorm.Open(postgres.Open(conString), &gorm.Config{})
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
