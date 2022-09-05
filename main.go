package main

import (
	"product-api/config"
	"product-api/database"
	"product-api/service"
)

func main() {
	DB := database.NewDB()

	h := service.NewEcho(*DB)
	config.MigrateProductTable()
	h.Logger.Fatal(h.Start(":1323"))
}
