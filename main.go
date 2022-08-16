package main

import (
	"product-api/database"
	"product-api/service"
)

func main() {
	DB := database.NewDB()
	h := service.NewEcho(*DB)
	h.Logger.Fatal(h.Start(":1323"))
}
