package controller

import (
	"product-api/model"
)

/*type DbRepo interface {
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Model(value interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
}*/
type ProductRepo interface {
	GetRepoProducts() ([]model.Product, error)
	AddNewProduct(newPro model.ProductRepository) (*model.Product, error)
	FindProductQueryParams(code, name, cate string) ([]model.Product, error)
	GetProductByID(id string) (model.Product, error)
	UpdateProductByID(mode model.ProductRepository) (string, error)
	DeleteProductByID(id string) (string, error)
}
