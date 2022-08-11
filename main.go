package main

import (
	"product-api/service"
)

func main() {
	//db := database.newDB()
	//rc := database.ConnectRedis()
	r := service.NewEcho()

	r.Logger.Fatal(r.Start("localhost:1324"))
}
