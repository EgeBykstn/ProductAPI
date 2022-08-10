package controller

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"product-api/model"
)

//GET
func GetProduct(c echo.Context) error {
	models := []model.Product{
		{Id: 1234, Code: "yönetici", Name: "ege", Category: "person", Price: 35, Color: "blue", Size: 180},
		{Id: 1234, Code: "yönetici", Name: "ege", Category: "person", Price: 55, Color: "blue", Size: 180},
		{Id: 1234, Code: "yönetici", Name: "ege", Category: "person", Price: 65, Color: "blue", Size: 180},
	}
	/*var products []model.Product

	err := pf.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	*/
	return c.JSON(http.StatusOK, models)
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
