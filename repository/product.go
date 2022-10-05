package repository

import (
	"gorm.io/gorm"
	"product-api/model"
	"time"
)

//go:generate mockgen -destination=../mocks/mock_repo.go -package=mocks product-api/repository GormRepo
type GormRepo interface {
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Model(value interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
}

type productRepo struct {
	location time.Location
	DB       GormRepo
}

func NewProductRepo(db GormRepo) *productRepo {
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

func (p *productRepo) AddNewRepoProduct(newPro *model.Product) error {
	newPro.CreatedAt = time.Now().In(&p.location)
	err := p.DB.Create(&newPro).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *productRepo) UpdateRepoProductByID(mode *model.Product) (string, error) {
	product := &model.Product{}
	product.UpdatedAt = time.Now().In(&p.location)

	err := p.DB.Where("id= ?", mode.Id).Find(&product).Error
	if err != nil {
		return "error", err
	}
	//err := p.DB.Model(&product).Updates(model.Product{
	//	Id:        mode.Id,
	//	Code:      mode.Code,
	//	Name:      mode.Name,
	//	Category:  mode.Category,
	//	Price:     mode.Price,
	//	Color:     mode.Color,
	//	Size:      mode.Size,
	//	UpdatedAt: time.Time{},
	//	CreatedAt: time.Time{},
	//}).Error
	/*	err = p.DB.First(&product).Error
		if err != nil {
			return "error", err
		}*/
	product = mode
	err = p.DB.Updates(product).Error
	if err != nil {
		return "error", err
	}
	return "product updated according to ID", nil
}

func (p *productRepo) FindRepoProductQueryParams(code, name, cate string) ([]model.Product, error) {
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
	//err := p.DB.Where(filter).Find(&Products).Error
	err := p.DB.Find(&Products, model.Product{Name: filter.Name, Code: filter.Code, Category: filter.Category}).Error
	if err != nil {
		return Products, err
	}

	return Products, nil
}

func (p *productRepo) GetRepoProductByID(id string) (*model.Product, error) {
	product := model.Product{}
	//err := p.DB.Where("id= ?", id).Find(&product).Error
	err := p.DB.First(&product, "id = ?", id).Error
	if err != nil {
		return &model.Product{}, err
	}

	return &product, nil
}

func (p *productRepo) DeleteRepoProductByID(id string) error {
	product := model.Product{}
	err := p.DB.Delete(product, id).Error
	if err != nil {
		return err
	}

	return err
}
