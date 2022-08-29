package controller

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"product-api/model"
	"strconv"
	"time"
)

type ProductController struct {
	DB gorm.DB
}
type Repository interface {
	GetProducts(c echo.Context) error
	AddNewProduct(c echo.Context) error
	FindProductQueryParams(c echo.Context) error
	GetProductByID(c echo.Context) error
	UpdateProductByID(c echo.Context) error
	DeleteProductByID(c echo.Context) error
}

func NewProductController(db gorm.DB) Repository {
	return &ProductController{
		DB: db,
	}
}

var _ Repository = ProductController{}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!!")
}

//GET
func (pc ProductController) GetProducts(c echo.Context) error {
	Products := []model.Product{}

	err := pc.DB.Find(&Products).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Products)
}

func (pc ProductController) GetProductByID(c echo.Context) error {
	Products := []model.Product{}
	ID := c.Param("id")
	StrID, _ := strconv.Atoi(ID)
	err := pc.DB.Find(&Products).Error
	if err != nil {
		return err
	}

	for i, p := range Products {
		if p.Id == StrID {
			return c.JSON(http.StatusOK, Products[i])
		}
	}
	return errors.New("product not found according to ID")
}
func (pc ProductController) FindProductQueryParams(c echo.Context) error {
	//filter := model.Product{}
	Pcode := c.QueryParam("code")
	Pname := c.QueryParam("name")
	var product []model.Product
	err := pc.DB.Where("code = ? AND name = ?", Pcode, Pname).Find(&product).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
	/*
		Products := []model.Product{}
		//filter_type := c.Param("type")
		Pcode := c.QueryParam("code")
		Pname := c.QueryParam("name")
		Pcate := c.QueryParam("category")
		Pprice, _ := strconv.Atoi(c.QueryParam("price"))
		Pcolor := c.QueryParam("color")
		Psize, _ := strconv.Atoi(c.QueryParam("size"))
		a := model.Filter{Code: Pcode, Name: Pname, Category: Pcate, Price: Pprice, Color: Pcolor, Size: Psize}
		err := pc.DB.Find(&Products).Error
		if err != nil {
			return c.JSON(http.StatusOK, "bulamadık")
		}
		for i, p := range Products {
			if p.Code == a.Code && p.Color == a.Color {
				return c.JSON(http.StatusOK, Products[i])
			}
		}

		return c.JSON(http.StatusOK, a)

	*/
}

//POST
func (pc ProductController) AddNewProduct(c echo.Context) error {
	NewProduct := new(model.Product)
	var ai model.AutoInc
	err := c.Bind(&NewProduct)
	if NewProduct.Name == "" || NewProduct.Color == "" || NewProduct.Code == "" || NewProduct.Category == "" || NewProduct.Size < 0 || NewProduct.Price < 0 {
		return c.JSON(http.StatusOK, "missing or wrong entered data")
	}

	NewProduct.CreatedAt = time.Now()
	NewProduct.UpdatedAt = time.Now()
	NewProduct.Id = ai.ID()
	pc.DB.Create(NewProduct)

	if err != nil {
		log.Printf("Failed Processing addInformation request")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed Processing addInformation request")
	}
	return c.JSONP(http.StatusOK, "Following Product Added \n", NewProduct)
}

func (pc ProductController) UpdateProductByID(c echo.Context) error {
	NewValueProduct := model.Product{}

	err := c.Bind(&NewValueProduct) // taking from body
	if err != nil {
		log.Printf("Failed Processing addInformation request")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed Processing addInformation request")
	}
	if NewValueProduct.Name == "" || NewValueProduct.Color == "" || NewValueProduct.Code == "" || NewValueProduct.Category == "" || NewValueProduct.Size < 0 || NewValueProduct.Price < 0 {
		return c.JSON(http.StatusOK, "missing or wrong entered data")
	}

	Products := []model.Product{}
	err = pc.DB.Find(&Products).Error
	if err != nil {
		return err
	}

	for _, p := range Products {
		if p.Id == NewValueProduct.Id {
			fmt.Println(p)

			p = NewValueProduct
			//	p.UpdatedAt = time.Now()
			pc.DB.Updates(p)

			//fmt.Println(model.Time{UpdatedAt: time.Now()})
			return c.JSON(http.StatusOK, p)
		}
	}

	return c.JSON(http.StatusOK, "not founded product according to ID")
}

func (pc ProductController) DeleteProductByID(c echo.Context) error {
	ID := c.Param("id")
	strID, _ := strconv.Atoi(ID)

	Products := []model.Product{}

	err := pc.DB.Find(&Products).Error
	if err != nil {
		return err
	}

	for _, p := range Products {
		if p.Id == strID {
			fmt.Println(p)

			pc.DB.Delete(p)
			return c.JSONP(http.StatusOK, "Product Deleted \n", p)
		}
	}

	return c.JSON(http.StatusOK, "not founded ID")
}
