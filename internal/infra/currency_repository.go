package infra

import (
	"github.com/lang-dev/exchange/internal/entity"
	"gorm.io/gorm"
)

type CurrencyInterface interface {
	Create(p entity.Currency) error
	FindByID(code string) (entity.Currency, error)
}

type CurrencyRepository struct {
	DB *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *CurrencyRepository {
	return &CurrencyRepository{
		DB: db,
	}
}

func (r *CurrencyRepository) Create(p entity.Currency) error {
	return r.DB.Create(p).Error
}

func (r *CurrencyRepository) FindByID(id string) (entity.Currency, error) {
	var currency entity.Currency
	err := r.DB.First(&currency, "id = ?", id).Error
	return currency, err
}
