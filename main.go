package main

import (
	"product-api/database"
	"product-api/service"
)

func main() {
	db := database.newDB()
	rc := database.ConnectRedis()
	r := service.NewEcho(*db, *rc)

	r.Logger.Fatal(r.Start("localhost:1324"))
}
