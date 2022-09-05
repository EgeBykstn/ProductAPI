package model

import (
	"product-api/database"
	"sync"
	"time"
)

type Product struct {
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

var Loc, _ = time.LoadLocation("Europe/Minsk")
var Products []Product

type AutoInc struct {
	sync.Mutex
	id int
}

/*func (s *AutoInc) ID() (id int) {
	s.Lock()
	defer s.Unlock()

	id = s.id
	s.id++
	return id
}*/

/*type ProductDBController struct {
	DB ProductRepository
}

func NewProductController(db ProductRepository) ProductDBController {
	return ProductDBController{
		DB: db,
	}
}*/

func (p *ProductRepository) GetRepoProducts() ([]Product, error) {
	db := database.GetDBInstance()
	err := db.Find(&Products).Error
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
