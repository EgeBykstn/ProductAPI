package controller

import (
	"product-api/model"
)

type ProductRepo interface {
	GetRepoProducts() ([]model.Product, error)
	AddNewProduct(newPro *model.Product) (*model.Product, error)
	FindProductQueryParams(code, name, cate string) ([]model.Product, error)
	GetRepoProductByID(id string) (*model.Product, error)
	UpdateProductByID(mode *model.Product) (string, error)
	DeleteProductByID(id string) error
}
