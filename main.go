package main

import (
	"product-api/database"
	"product-api/service"
	"time"
)

func main() {
	<-time.After(time.Second * 4)
	DB := database.NewDB()
	h := service.NewEcho(DB)
	database.MigrateProductTable(DB)
	h.Logger.Fatal(h.Start(":1323"))
}
