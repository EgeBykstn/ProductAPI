package controller

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"product-api/model"
)

//GET
func (pf *ProductController) GetProduct() ([]model.Product, error) {
	var products []model.Product

	err := pf.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

//POST
func addInfo(c echo.Context) error {
	product := model.Product{}

	err := c.Bind(&product)

	if err != nil {
		log.Printf("Failed Processing addInformation request")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed Processing addInformation request")
	}
	log.Printf("this is your informations: ", product)
	return c.String(http.StatusOK, "WE GOT YOUR INFORMATION	")

}
