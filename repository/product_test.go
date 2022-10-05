package repository

import (
	"fmt"
	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"math/rand"
	"product-api/mocks"
	"product-api/model"
	strconv "strconv"
	"testing"
	"time"
)

var productRep *productRepo
var mockGormRepo *mocks.MockGormRepo

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGormRepo = mocks.NewMockGormRepo(ctrl)
	productRep = NewProductRepo(mockGormRepo)

	return func() {
		productRep = nil

		defer ctrl.Finish()
	}
}
func TestGetProducts_Success(t *testing.T) {
	td := setup(t)
	defer td()

	Response := &gorm.DB{
		Error: nil,
	}
	mockRows := []model.Product{RandomProduct()}

	mockGormRepo.EXPECT().Find(gomock.Any()).SetArg(0, mockRows).Return(Response)
	products, err := productRep.GetRepoProducts()

	require.NoError(t, err)
	require.NotNil(t, products)
}

func TestProductRepo_GetRepoProducts_Unsuccessful(t *testing.T) {
	td := setup(t)
	defer td()

	err := fmt.Errorf("error expecting")

	Response := &gorm.DB{
		Error: err,
	}
	mockGormRepo.EXPECT().Find(gomock.Any()).Return(Response)
	result, err := productRep.GetRepoProducts()

	require.Error(t, err)
	require.Nil(t, result)

}

func TestAddNewRepoProduct_Success(t *testing.T) {
	td := setup(t)
	defer td()

	Response := &gorm.DB{
		Error: nil,
	}
	Product := RandomProduct()
	mockGormRepo.EXPECT().Create(gomock.Any()).Return(Response)
	err := productRep.AddNewRepoProduct(&Product)

	require.NoError(t, err)
	//require.NotNil(t, NewProduct)
	//require.Equal(t, NewProduct, Product)
}
func TestAddNewRepoProduct_Unsuccessful(t *testing.T) {
	td := setup(t)
	defer td()

	err := fmt.Errorf("error expecting")

	Response := &gorm.DB{
		Error: err,
	}
	Product := RandomProduct()
	mockGormRepo.EXPECT().Create(gomock.Any()).Return(Response)
	err = productRep.AddNewRepoProduct(&Product)

	require.Error(t, err)
	//require.NotNil(t, NewProduct)
	//require.Equal(t, NewProduct, Product)
}

/*func TestUpdateRepoProductByID_Success(t *testing.T) {
	td := setup(t)
	defer td()

	Response := &gorm.DB{
		Error: nil,
	}

	a := RandomProduct()
	err := mockGormRepo.Where(gomock.Any(), gomock.Any()).Find(gomock.Any()).Error
	if err != nil {
		require.Error(t, err)
	}
	mockGormRepo.EXPECT().Updates(gomock.Any()).Return(Response)

	result, err := productRep.UpdateRepoProductByID(&a)

	require.NoError(t, err)
	require.NotNil(t, result)
}*/

/*func TestUpdateRepoProductByID_Unsuccessful(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	err := fmt.Errorf("error")
	mockGormRepo := mocks.NewMockGormRepo(ctrl)
	productRepo := NewProductRepo(mockGormRepo)
	Response := &gorm.DB{
		Error: err,
	}

	a := RandomProduct()
	mockGormRepo.EXPECT().First(gomock.Any()).Return(Response)
	mockGormRepo.EXPECT().Updates(gomock.Any()).Return(Response)

	result, err := productRepo.UpdateRepoProductByID(&a)

	require.Error(t, err)
	require.NotNil(t, result)
}*/
func TestGetRepoProductByID_Success(t *testing.T) {
	td := setup(t)
	defer td()

	Response := &gorm.DB{
		Error: nil,
	}

	mockRows := []model.Product{RandomProduct()}
	fakeID := strconv.Itoa(mockRows[0].Id)

	//mockGormRepo.EXPECT().Where(gomock.Any(), gomock.Any()).Return(Response)
	//mockGormRepo.EXPECT().Find(gomock.Any()).Return(Response)
	mockGormRepo.EXPECT().First(gomock.Any(), gomock.Any()).Return(Response)

	result, err := productRep.GetRepoProductByID(fakeID)

	require.NoError(t, err)
	require.NotNil(t, result)

}
func TestGetRepoProductByID_Unsuccessful(t *testing.T) {
	td := setup(t)
	defer td()

	err := fmt.Errorf("error")

	Response := &gorm.DB{
		Error: err,
	}

	mockRows := []model.Product{RandomProduct()}
	fakeID := strconv.Itoa(mockRows[0].Id)

	//mockGormRepo.EXPECT().Where(gomock.Any(), gomock.Any()).Return(Response)
	//mockGormRepo.EXPECT().Find(gomock.Any()).Return(Response)
	mockGormRepo.EXPECT().First(gomock.Any(), gomock.Any()).Return(Response)

	result, err := productRep.GetRepoProductByID(fakeID)

	require.Error(t, err)
	require.NotNil(t, result)

}
func TestFindRepoProductQueryParams(t *testing.T) {
	td := setup(t)
	defer td()

	response := &gorm.DB{
		Error: nil,
	}
	mocksRows := []model.Product{RandomProduct()}
	//random := RandomProduct()
	mockGormRepo.EXPECT().Find(gomock.Any(), gomock.Any()).Return(response)

	result, err := productRep.FindRepoProductQueryParams(mocksRows[0].Code, mocksRows[0].Name, mocksRows[0].Category)

	require.NoError(t, err)
	require.Nil(t, result)

}
func TestFindRepoProductQueryParams_Unsuccessful(t *testing.T) {
	td := setup(t)
	defer td()

	err := fmt.Errorf("error")

	response := &gorm.DB{
		Error: err,
	}
	mocksRows := []model.Product{RandomProduct()}
	//random := RandomProduct()
	mockGormRepo.EXPECT().Find(gomock.Any(), gomock.Any()).Return(response)

	result, err := productRep.FindRepoProductQueryParams(mocksRows[0].Code, mocksRows[0].Name, mocksRows[0].Category)

	require.Error(t, err)
	require.Nil(t, result)

}
func TestDeleteRepoProductByID_Success(t *testing.T) {
	td := setup(t)
	defer td()

	Response := &gorm.DB{
		Error: nil,
	}

	mockRows := []model.Product{RandomProduct()}
	fakeID := strconv.Itoa(mockRows[0].Id)

	mockGormRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(Response)
	err := productRep.DeleteRepoProductByID(fakeID)
	require.NoError(t, err)
}
func TestDeleteRepoProductByID_Unsuccessful(t *testing.T) {
	td := setup(t)
	defer td()

	err := fmt.Errorf("error")

	Response := &gorm.DB{
		Error: err,
	}

	mockRows := []model.Product{RandomProduct()}
	fakeID := strconv.Itoa(mockRows[0].Id)

	mockGormRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(Response)
	err = productRep.DeleteRepoProductByID(fakeID)
	require.Error(t, err)
}

func RandomProduct() model.Product {
	return model.Product{
		Id:        rand.Intn(10),
		Code:      random.String(4),
		Name:      random.String(4),
		Category:  random.String(4),
		Price:     rand.Intn(100),
		Color:     random.String(4),
		Size:      rand.Intn(100),
		UpdatedAt: time.Time{},
		CreatedAt: time.Time{},
	}
}
