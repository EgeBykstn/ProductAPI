package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"product-api/model"
)

type productController struct {
	ProductRepository ProductRepo
}

func NewProductController(productDBRepository ProductRepo) *productController {
	return &productController{
		ProductRepository: productDBRepository,
	}
}

//GetProducts list all products from product table as an array
func (pc productController) GetProducts(c echo.Context) error {
	products, err := pc.ProductRepository.GetRepoProducts()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}

//GetProductByID gets product from product table according to selected ID
func (pc productController) GetProductByID(c echo.Context) error {
	ID := c.Param("id")
	product, err := pc.ProductRepository.GetRepoProductByID(ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}

//FindProductQueryParams searches for product according to query parameters
func (pc productController) FindProductQueryParams(c echo.Context) error {
	pCode := c.QueryParam("code")
	pName := c.QueryParam("name")
	pCate := c.QueryParam("category")
	product, err := pc.ProductRepository.FindRepoProductQueryParams(pCode, pName, pCate)
	if err != nil {
		return err
	}
	if product == nil {
		return c.JSON(http.StatusOK, "there is no product according to entered data")
	}

	return c.JSON(http.StatusOK, product)
}

//AddNewProduct adds new product with entered data to product table
func (pc productController) AddNewProduct(c echo.Context) error {
	NewProduct := model.Product{}
	err := c.Bind(&NewProduct)
	NewProduct.Id = 0
	if NewProduct.Name == "" || NewProduct.Color == "" || NewProduct.Code == "" || NewProduct.Category == "" || NewProduct.Size < 0 || NewProduct.Price < 0 {
		return c.JSON(http.StatusOK, "missing or wrong entered data")
	}
	err = pc.ProductRepository.AddNewRepoProduct(&NewProduct)
	if err != nil {
		return err
	}

	return c.JSONP(http.StatusOK, "Following Product Added %\n", err)
}

//UpdateProductByID updates selected product according to ID
func (pc productController) UpdateProductByID(c echo.Context) error {
	NewValueProduct := model.Product{}
	err := c.Bind(&NewValueProduct)
	/*if err != nil {
		return err
	}*/
	if NewValueProduct.Name == "" || NewValueProduct.Color == "" || NewValueProduct.Code == "" || NewValueProduct.Category == "" || NewValueProduct.Size < 0 || NewValueProduct.Price < 0 {
		return c.JSON(http.StatusOK, "missing or wrong entered data")
	}
	UpdatedProduct, _ := pc.ProductRepository.UpdateRepoProductByID(&NewValueProduct)
	if err != nil {
		return c.JSON(http.StatusOK, "type error")
	}

	return c.JSON(http.StatusOK, UpdatedProduct)
}

//DeleteProductByID deletes product according to selected ID
func (pc productController) DeleteProductByID(c echo.Context) error {
	ID := c.Param("id")
	deletedPro := pc.ProductRepository.DeleteRepoProductByID(ID)

	return c.JSON(http.StatusOK, deletedPro)
}
