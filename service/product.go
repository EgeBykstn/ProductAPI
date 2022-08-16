package service

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"product-api/controller"
)

type handler struct {
	DB *gorm.DB
}

func NewEcho(db gorm.DB) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", controller.Hello)
	e.GET("/products", controller.GetProduct)
	e.POST("/products", controller.AddNewProduct)
	e.GET("/products/:id", controller.ProductByID)
	e.PUT("/products/:id", controller.UpdateProductByID)
	e.DELETE("/delete-product/:id", controller.DeleteProductByID)
	/*
		e.GET("/search", h.FindProduct)
		e.GET("/search-params", h.FindProductQueryParams)

		e.POST("/batch-products", h.BatchCreateProduct)

	*/
	return e
}
