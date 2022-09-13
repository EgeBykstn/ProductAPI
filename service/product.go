package service

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"product-api/controller"
	"product-api/handler"
)

func NewEcho(db *gorm.DB) *echo.Echo {
	e := echo.New()
	productRepo := handler.NewProductRepo(db)
	pc := controller.NewProductController(productRepo)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/products", pc.GetProducts)
	e.GET("/filter-products", pc.FindProductQueryParams)
	e.POST("/products", pc.AddNewProduct)
	e.GET("/products/:id", pc.GetProductByID)
	e.PUT("/update-product", pc.UpdateProductByID)
	e.DELETE("/delete-product/:id", pc.DeleteProductByID)

	return e
}
