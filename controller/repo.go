package controller

import (
	"product-api/model"
)

//go:generate mockgen -destination=../mocks/mock_controller.go -package=mocks product-api/controller ProductRepo
type ProductRepo interface {
	GetRepoProducts() ([]model.Product, error)
	AddNewRepoProduct(newPro *model.Product) error
	FindRepoProductQueryParams(code, name, cate string) ([]model.Product, error)
	GetRepoProductByID(id string) (*model.Product, error)
	UpdateRepoProductByID(mode *model.Product) (string, error)
	DeleteRepoProductByID(id string) error
}
