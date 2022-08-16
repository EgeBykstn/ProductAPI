package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"product-api/database"
	"product-api/model"
	"strconv"
)

type Handler struct {
	DB *gorm.DB
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!!")
}

//GET
func GetProduct(c echo.Context) error {
	products, _ := GetRepoProducts()
	return c.JSON(http.StatusOK, products)
}
func GetRepoProducts() ([]model.Product, error) {
	db := database.GetDBInstance()
	Products := []model.Product{}

	if err := db.Find(&Products).Error; err != nil {
		return nil, err
	}

	return Products, nil
}
func ProductByID(c echo.Context) error {
	ID := c.Param("id")
	s2, _ := strconv.Atoi(ID)
	product, err := GetRepoProductByID(s2)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product Not Found")
	}
	return c.JSON(http.StatusOK, product)
}
func GetRepoProductByID(id int) (*model.Product, error) {
	db := database.GetDBInstance()
	Products := []model.Product{}
	if err := db.Find(&Products).Error; err != nil {
		return nil, err
	}

	for i, p := range Products {
		if p.Id == id {
			return &Products[i], nil
		}
	}
	return nil, errors.New("product not found according to ID")
}

//POST
func AddNewProduct(c echo.Context) error {
	NewProduct := new(model.Product)
	err := c.Bind(&NewProduct)
	db := database.GetDBInstance()
	db.Create(NewProduct)
	if err != nil {
		log.Printf("Failed Processing addInformation request")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed Processing addInformation request")
	}

	log.Printf("this is your informations: ", NewProduct)
	return c.JSON(http.StatusOK, NewProduct)
}
func UpdateProductByID(c echo.Context) error {
	UpdatedProduct := new(model.Product)
	err := c.Bind(&UpdatedProduct)
	ID := c.Param("id")
	s2, _ := strconv.Atoi(ID)
	product, err := GetRepoProductByID(s2)

	product = UpdatedProduct

	if err != nil {
		return c.JSON(http.StatusNotFound, "Product Not Found")
	}
	db := database.GetDBInstance()
	db.Save(product)
	return c.JSON(http.StatusOK, "Products Updated according to ID")
}
func DeleteProductByID(c echo.Context) error {
	ID := c.Param("id")
	s2, _ := strconv.Atoi(ID)
	product, err := GetRepoProductByID(s2)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product Not Found")
	}
	db := database.GetDBInstance()
	db.Delete(product)

	return c.JSON(http.StatusOK, "Following Product Deleted")
}
