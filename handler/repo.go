package handler

import (
	"fmt"
	"product-api/database"
	"product-api/model"
	"time"
)

type productRepo struct {
	location time.Location
	DB       database.GormDB
}

func NewProductRepo(db database.GormDB) *productRepo {
	return &productRepo{
		DB: db,
	}
}

func (p *productRepo) GetRepoProducts() ([]model.Product, error) {
	var Products []model.Product
	err := p.DB.Find(&Products).Error
	if err != nil {
		return Products, err
	}

	return Products, nil
}

func (p *productRepo) AddNewProduct(newPro *model.Product) (*model.Product, error) {
	newPro.CreatedAt = time.Now().In(&p.location)
	err := p.DB.Create(&newPro).Error
	if err != nil {
		return &model.Product{}, err
	}
	return newPro, nil
}

func (p *productRepo) UpdateProductByID(mode *model.Product) (string, error) {
	product := &model.Product{}
	err := p.DB.Where("id= ?", mode.Id).Find(&product).Error
	if err != nil {
		return "no match with selected ID", err
	}
	product = mode
	product.UpdatedAt = time.Now().In(&p.location)
	p.DB.Updates(product)

	return "product updated according to ID", nil
}

func (p *productRepo) FindProductQueryParams(code, name, cate string) ([]model.Product, error) {
	var Products []model.Product
	filter := model.Product{}
	if code != "" {
		filter.Code = code
	}
	if name != "" {
		filter.Name = name
	}
	if cate != "" {
		filter.Category = cate
	}

	err := p.DB.Where(filter).Find(&Products).Error
	if err != nil {
		return Products, err
	}

	return Products, nil
}

func (p *productRepo) GetRepoProductByID(id string) (*model.Product, error) {
	product := model.Product{}
	err := p.DB.Where("id= ?", id).Find(&product).Error
	if err != nil {
		return &model.Product{}, err
	}

	return &product, nil
}

func (p *productRepo) DeleteProductByID(id string) error {
	product := model.Product{}
	err := p.DB.Where("id= ?", id).Find(&product).Error
	if err != nil {
		return fmt.Errorf("find error: %v", err)

	}
	if product.Id == 0 {
		return err
	}

	err = p.DB.Delete(product).Error
	if err != nil {
		return err
	}

	return fmt.Errorf("deleted: %v", err)
}
