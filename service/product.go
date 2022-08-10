package service

import (
	"github.com/labstack/echo/v4"
	"product-api/controller"
)

func NewEcho() *echo.Echo {
	e := echo.New()
	/*	pc := controller.NewProductController(&db, rc)
		pf := http.Repo(pc)
		h := http.NewHandler(pf)

		e.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	*/

	e.GET("/products", controller.GetProduct)
	e.POST("/products", controller.AddNewProduct)
	e.GET("/products/:id", controller.ProductByID)
	/*e.GET("/products/:id", h.GetProductByID)
	e.GET("/search", h.FindProduct)
	e.GET("/search-params", h.FindProductQueryParams)
	e.POST("/products", h.CreateProduct)
	e.POST("/batch-products", h.BatchCreateProduct)
	e.PUT("/products/:id", h.UpdateProduct)
	e.DELETE("/products/:id", h.DeleteProduct)
	*/
	return e
}
