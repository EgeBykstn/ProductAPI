package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"product-api/model"
	"strconv"
)

var products = []model.Product{
	{Id: 1234, Code: "yönetici", Name: "ege", Category: "person", Price: 35, Color: "blue", Size: 180},
	{Id: 1134, Code: "yönetici", Name: "ege", Category: "person", Price: 55, Color: "blue", Size: 180},
	{Id: 1334, Code: "yönetici", Name: "ege", Category: "person", Price: 65, Color: "blue", Size: 180},
}

//GET
func GetProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

//POST
func AddNewProduct(c echo.Context) error {
	var newProduct model.Product
	err := c.Bind(&newProduct)

	if err != nil {
		log.Printf("Failed Processing addInformation request")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed Processing addInformation request")
	}
	products = append(products, newProduct)
	log.Printf("this is your informations: ", products)
	return c.JSON(http.StatusOK, newProduct)

}
func ProductByID(c echo.Context) error {
	ID := c.QueryParam("id")
	s2, err := strconv.Atoi(ID)
	product, err := GetProductByID(s2)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, product)

}
func GetProductByID(id int) (*model.Product, error) {
	for i, p := range products {
		if p.Id == id {
			return &products[i], nil
		}
	}
	return nil, errors.New("product not found according to ID")
}
