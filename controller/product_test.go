package controller

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"product-api/mocks"
	"product-api/model"
	"testing"
	"time"
)

var productCont *productController
var mockProductRepo *mocks.MockProductRepo

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo = mocks.NewMockProductRepo(ctrl)
	productCont = NewProductController(mockProductRepo)

	return func() {
		productCont = nil
		defer ctrl.Finish()
	}
}

func TestGetProducts_Success(t *testing.T) {
	td := setup(t)
	defer td()

	e := echo.New()
	e.GET("/products", productCont.GetProducts)

	var FakeDataForTest = []model.Product{
		{123, "007", "123", "123", 123, "123", 123, time.Now(), time.Now()},
		{321, "001", "123", "123", 123, "123", 123, time.Now(), time.Now()},
		{4213, "790", "123", "123", 123, "123", 123, time.Now(), time.Now()},
	}

	mockProductRepo.EXPECT().GetRepoProducts().Return(FakeDataForTest, nil)

	request := httptest.NewRequest("GET", "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(request, rec)
	// Assertions
	if assert.NoError(t, productCont.GetProducts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//assert.Equal(t, userJSON, rec.Body.String())
	}
}
func TestGetProducts_Unsuccessful(t *testing.T) {
	td := setup(t)
	defer td()

	e := echo.New()
	e.GET("/products", productCont.GetProducts)

	var FakeDataForTest = []model.Product{
		{123, "007", "123", "123", 123, "123", 123, time.Now(), time.Now()},
		{321, "001", "123", "123", 123, "123", 123, time.Now(), time.Now()},
		{4213, "790", "123", "123", 123, "123", 123, time.Now(), time.Now()},
	}
	err := error(fmt.Errorf("TestError"))
	mockProductRepo.EXPECT().GetRepoProducts().Return(FakeDataForTest, err)

	request := httptest.NewRequest("GET", "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(request, rec)
	// Assertions
	if assert.Error(t, productCont.GetProducts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetProductByID_Success(t *testing.T) {
	td := setup(t)
	defer td()

	e := echo.New()
	e.GET("/products/:id", productCont.GetProductByID)

	FakeDataForTest := RandomProduct()
	mockProductRepo.EXPECT().GetRepoProductByID(gomock.Any()).Return(&FakeDataForTest, nil)

	req := httptest.NewRequest("GET", "/products/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := productCont.GetProductByID(c)

	require.NoError(t, err)

}
func TestGetProductByID_Unsuccessful(t *testing.T) {
	td := setup(t)
	defer td()

	e := echo.New()
	e.GET("/products/:id", productCont.GetProductByID)
	err := fmt.Errorf("TestError")
	FakeDataForTest := RandomProduct()
	mockProductRepo.EXPECT().GetRepoProductByID(gomock.Any()).Return(&FakeDataForTest, err)

	req := httptest.NewRequest("GET", "/products/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = productCont.GetProductByID(c)

	require.Error(t, err)
	require.Equal(t, http.StatusOK, rec.Code)

}
func TestFindProductQueryParams_Success(t *testing.T) {
	td := setup(t)
	defer td()

	var array []model.Product
	e := echo.New()
	e.GET("/find-products", productCont.FindProductQueryParams)
	mockProductRepo.EXPECT().FindRepoProductQueryParams(gomock.Any(), gomock.Any(), gomock.Any()).Return(array, nil)

	req := httptest.NewRequest("GET", "/find-products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := productCont.FindProductQueryParams(c)
	require.NoError(t, err)

}
func TestFindProductQueryParams_Unsuccessful(t *testing.T) {
	td := setup(t)
	defer td()
	//FakeDataForTest := RandomProduct()
	var array []model.Product
	err := fmt.Errorf("TestError")
	e := echo.New()
	e.GET("/find-products", productCont.FindProductQueryParams)
	mockProductRepo.EXPECT().FindRepoProductQueryParams(gomock.Any(), gomock.Any(), gomock.Any()).Return(array, err)

	req := httptest.NewRequest("GET", "/find-products", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	err = productCont.FindProductQueryParams(c)
	require.Error(t, err)
	//require.NotNil(t, array)
	require.Equal(t, http.StatusOK, res.Code)

}

func TestAddNewProduct_Success(t *testing.T) {
	td := setup(t)
	defer td()

	e := echo.New()
	e.POST("/products", productCont.AddNewProduct)
	mockProductRepo.EXPECT().AddNewRepoProduct(gomock.Any()).Return(nil)

	req := httptest.NewRequest("POST", "/products", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	err := productCont.AddNewProduct(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.Code)
}
func TestAddNewProduct_Unsuccessful(t *testing.T) {
	td := setup(t)
	defer td()

	err := fmt.Errorf("TestError")
	e := echo.New()

	mockProductRepo.EXPECT().AddNewRepoProduct(gomock.Any()).Return(err)
	req := httptest.NewRequest("POST", "/products", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	err = productCont.AddNewProduct(c)

	require.Nil(t, err)
	require.Equal(t, http.StatusOK, res.Code)

}

func TestUpdateProductByID_Success(t *testing.T) {
	td := setup(t)
	defer td()

	FakeDataForTest := RandomProduct()
	e := echo.New()
	e.PUT("/update-product", productCont.UpdateProductByID)
	mockProductRepo.EXPECT().UpdateRepoProductByID(&FakeDataForTest).Return(RandomProduct().Code, nil)

	req := httptest.NewRequest("PUT", "/update-products", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	err := productCont.UpdateProductByID(c)
	require.NoError(t, err)

}
func TestUpdateProductByID_Unsuccessful(t *testing.T) {
	td := setup(t)
	defer td()

	FakeDataForTest := RandomProduct()
	err1 := fmt.Errorf("TestError")
	e := echo.New()
	e.PUT("/update-product", productCont.UpdateProductByID)
	mockProductRepo.EXPECT().UpdateRepoProductByID(&FakeDataForTest).Return(RandomProduct().Code, err1)

	req := httptest.NewRequest("PUT", "/update-products", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	require.Error(t, err1)
	err := productCont.UpdateProductByID(c)
	require.NoError(t, err)

}
func TestDeleteProductByID_Success(t *testing.T) {
	td := setup(t)
	defer td()

	e := echo.New()
	e.DELETE("/delete-product/:id", productCont.DeleteProductByID)
	mockProductRepo.EXPECT().DeleteRepoProductByID(gomock.Any()).Return(nil)

	req := httptest.NewRequest("DELETE", "/delete-product/:id", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	err := productCont.DeleteProductByID(c)
	require.NoError(t, err)

}
func TestDeleteProductByID_Unsuccessful(t *testing.T) {
	td := setup(t)
	defer td()

	err := fmt.Errorf("TestError")
	e := echo.New()
	e.DELETE("/delete-product/:id", productCont.DeleteProductByID)
	mockProductRepo.EXPECT().DeleteRepoProductByID(gomock.Any()).Return(err)

	req := httptest.NewRequest("DELETE", "/delete-product/:id", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	err2 := productCont.DeleteProductByID(c)
	require.Error(t, err)
	require.NoError(t, err2)
}

func RandomProduct() model.Product {
	return model.Product{
		Id:        rand.Int(),
		Code:      random.String(4),
		Name:      random.String(4),
		Category:  random.String(4),
		Price:     rand.Int(),
		Color:     random.String(4),
		Size:      rand.Int(),
		UpdatedAt: time.Time{},
		CreatedAt: time.Time{},
	}

}
