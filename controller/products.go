package controller

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"product-api/model"
	"time"
)

type ProductRepoController struct {
	DB ProductRepo
}

func NewProductController(db ProductRepo) ProductRepoController {
	return ProductRepoController{
		DB: db,
	}
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!!")
}

//GET
func (pc ProductRepoController) GetProducts(c echo.Context) error {
	products, err := pc.DB.GetRepoProducts()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}

func (pc ProductRepoController) GetProductByID(c echo.Context) error {
	ID := c.Param("id")
	product, err := pc.DB.GetProductByID(ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, product)
}
func (pc ProductRepoController) FindProductQueryParams(c echo.Context) error {
	pCode := c.QueryParam("code")
	pName := c.QueryParam("name")
	pCate := c.QueryParam("category")

	product, err := pc.DB.FindProductQueryParams(pCode, pName, pCate)

	if err != nil {
		return err
	}
	if product == nil {
		return c.JSON(http.StatusOK, "there is no product according to entered data")
	}
	return c.JSON(http.StatusOK, product)
}

//POST
func (pc ProductRepoController) AddNewProduct(c echo.Context) error {
	NewProduct := model.ProductRepository{}
	//	var ai model.AutoInc
	err := c.Bind(&NewProduct)
	if err != nil {
		return c.JSON(http.StatusOK, "wrong data type")
	}
	NewProduct.Id = 0
	if NewProduct.Name == "" || NewProduct.Color == "" || NewProduct.Code == "" || NewProduct.Category == "" || NewProduct.Size < 0 || NewProduct.Price < 0 {
		return c.JSON(http.StatusOK, "missing or wrong entered data")
	}
	//Loc, _ := time.LoadLocation("Europe/Minsk")
	NewProduct.CreatedAt = time.Now()
	//NewProduct.Id = ai.ID()

	lastProduct, err := pc.DB.AddNewProduct(NewProduct)
	if err != nil {
		return c.JSON(http.StatusOK, "error")
	}
	time.Now().Local()
	return c.JSONP(http.StatusOK, "Following Product Added \n", lastProduct)
}

func (pc ProductRepoController) UpdateProductByID(c echo.Context) error {
	NewValueProduct := model.ProductRepository{}

	err := c.Bind(&NewValueProduct)
	if err != nil {
		log.Printf("Failed Processing addInformation request")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed Processing addInformation request")
	}
	if NewValueProduct.Name == "" || NewValueProduct.Color == "" || NewValueProduct.Code == "" || NewValueProduct.Category == "" || NewValueProduct.Size < 0 || NewValueProduct.Price < 0 {
		return c.JSON(http.StatusOK, "missing or wrong entered data")
	}
	UpdatedProduct, err := pc.DB.UpdateProductByID(NewValueProduct)
	NewValueProduct.UpdatedAt = time.Now().In(model.Loc)
	if err != nil {
		return c.JSON(http.StatusOK, "type error")
	}
	return c.JSON(http.StatusOK, UpdatedProduct)
}

func (pc ProductRepoController) DeleteProductByID(c echo.Context) error {
	ID := c.Param("id")
	deletedPro, err := pc.DB.DeleteProductByID(ID)
	if err != nil {
		return c.JSON(http.StatusOK, "")
	}
	return c.JSON(http.StatusOK, deletedPro)
}
