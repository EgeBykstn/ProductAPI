package model

import (
	"gorm.io/gorm"
	"product-api/database"
	"time"
)

type ProductRepository struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Price     int       `json:"price"`
	Color     string    `json:"color"`
	Size      int       `json:"size"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time
}

type DbRepo interface {
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Model(value interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
}
type ProductController struct {
	Db DbRepo
}

func NewProductController(db DbRepo) *ProductController {
	return &ProductController{
		Db: db,
	}
}

func (p *ProductRepository) GetRepoProducts() ([]Product, error) {
	//db := database.GetDBInstance()
	err := DbRepo.Find(&Products).Error
	if err != nil {
		return Products, err
	}

	return Products, nil
}
func (p *ProductRepository) AddNewProduct(newPro ProductRepository) (*Product, error) {
	db := database.GetDBInstance()
	NewProduct := Product{}
	NewProduct = Product(newPro)
	NewProduct.CreatedAt = time.Now().In(Loc)
	err := db.Create(&NewProduct).Error
	if err != nil {
		return &Product{}, err
	}
	return &NewProduct, nil
}
func (p *ProductRepository) UpdateProductByID(mode ProductRepository) (string, error) {
	product := Product{}
	db := database.GetDBInstance()
	err := db.Where("id= ?", mode.Id).Find(&product).Error
	if err != nil {
		return "no match with selected ID", err
	}
	product = Product(mode)
	db.Updates(product)
	return "product updated according to ID", nil
}
func (p *ProductRepository) FindProductQueryParams(code, name, cate string) ([]Product, error) {
	db := database.GetDBInstance()

	filter := ProductRepository{}
	if code != "" {
		filter.Code = code
	}
	if name != "" {
		filter.Name = name
	}
	if cate != "" {
		filter.Category = cate
	}

	err := db.Where(filter).Find(&Products).Error
	if err != nil {
		return Products, err
	}

	return Products, nil
}
func (p *ProductRepository) GetProductByID(id string) (Product, error) {
	product := Product{}
	db := database.GetDBInstance()
	err := db.Where("id= ?", id).Find(&product).Error
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (p *ProductRepository) DeleteProductByID(id string) (string, error) {
	product := Product{}
	db := database.GetDBInstance()
	err := db.Where("id= ?", id).Find(&product).Error
	if err != nil {
		return "error", err
	}
	if product.Id == 0 {
		return "no product match with ID", err
	}
	db.Delete(product)
	return "product deleted according to ID", nil
}
