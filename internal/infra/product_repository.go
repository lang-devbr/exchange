package infra

import (
	"github.com/lang-dev/exchange/internal/entity"
	"gorm.io/gorm"
)

type ProductInterface interface {
	Create(p entity.Product) error
	FindByID(id string) (entity.Product, error)
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) Create(p entity.Product) error {
	return r.DB.Create(p).Error
}

func (r *ProductRepository) FindByID(id string) (entity.Product, error) {
	var product entity.Product
	err := r.DB.First(&product, "id = ?", id).Error
	return product, err
}

func  GetAll (db *gorm.DB) ([]entity.Product, error ) { 
	var products []entity.Product 
	err := db.Model(&entity.Product{}).Preload( "currencies" ).Find(&products).Error
	return products, err
}
    

    

