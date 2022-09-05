package config

import (
	"encoding/json"
	"fmt"
	"os"
	"product-api/database"
	"product-api/model"
)

type Config struct {
	DBUser     string `json:"DBUser"`
	DBPassword string `json:"DBPassword"`
	DBName     string `json:"DBName"`
	DBHost     string `json:"DBHost"`
	DBPort     string `json:"DBPort"`
	DBType     string `json:"DBType"`
	IsMigrate  bool   `json:"isMigrate"`
}

func MigrateProductTable() {
	DB := database.GetDBInstance()
	c := Config{}
	if c.IsMigrate == false {
		err := DB.AutoMigrate(&model.Product{})
		if err != nil {
			panic("migration failed")
		}
	}
}

func GetConnString() Config {
	var cfg Config

	cfgFile, _ := os.Open("config/config.json")
	defer cfgFile.Close()

	jsonParser := json.NewDecoder(cfgFile)
	jsonParser.Decode(&cfg)
	return cfg
}
func GetDBType() string {
	return GetConnString().DBType
}

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		GetConnString().DBHost,
		GetConnString().DBPort,
		GetConnString().DBUser,
		GetConnString().DBName,
		GetConnString().DBPassword)
	return dataBase
}
