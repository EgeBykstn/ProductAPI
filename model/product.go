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
	DB controller.DbRepo
}

func NewProductController(db controller.DbRepo) ProductDBController {
	return ProductDBController{
		DB: db,
	}
}
func (p *ProductDBController) GetRepoProducts() ([]Product, error) {
	//db := database.GetDBInstance()
	err := p.DB.Find(&Products).Error
	if err != nil {
		return Products, err
	}

	return Products, nil
}*/

func (p *Product) GetRepoProducts() ([]Product, error) {
	db := database.GetDBInstance()
	err := db.Find(&Products).Error
	if err != nil {
		return Products, err
	}

	return Products, nil
}
func (p *Product) AddNewProduct(newPro Product) (*Product, error) {
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
func (p *Product) UpdateProductByID(mode Product) (string, error) {
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
func (p *Product) FindProductQueryParams(code, name, cate string) ([]Product, error) {
	db := database.GetDBInstance()

	filter := Product{}
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
func (p *Product) GetProductByID(id string) (Product, error) {
	product := Product{}
	db := database.GetDBInstance()
	err := db.Where("id= ?", id).Find(&product).Error
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (p *Product) DeleteProductByID(id string) (string, error) {
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
