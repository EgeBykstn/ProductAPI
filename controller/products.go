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
	{Id: 1, Code: "yönetici", Name: "ege", Category: "person", Price: 135, Color: "blue", Size: 180},
	{Id: 2, Code: "çalışan", Name: "ege", Category: "person", Price: 55, Color: "blue", Size: 180},
	{Id: 3, Code: "head", Name: "ege", Category: "person", Price: 65, Color: "blue", Size: 180},
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
	ID := c.Param("id")
	s2, _ := strconv.Atoi(ID)
	product, err := GetProductByID(s2)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product Not Found")
	}
	return c.JSON(http.StatusOK, product)
}

//sürekli nor found veriyor
func GetProductByID(id int) (*model.Product, error) {
	for i, p := range products {
		if p.Id == id {
			return &products[i], nil
		}
	}
	return nil, errors.New("product not found according to ID")
}

/*
func UpdateAvailableProduct(c echo.Context) error {
	var newProduct model.Product
	err := c.Bind(&newProduct)
	if err != nil {
		log.Printf("Failed Processing addInformation request")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed Processing addInformation request")
	}
	ID := c.Param("id")
	s2, _ := strconv.Atoi(ID)
	product, err := GetProductByID(s2)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product Not Found")
	}
	products.
	model.Product(&product
	product = &newProduct
	log.Printf("this is your informations: ", products)
	return c.JSON(http.StatusOK, newProduct)
} */
