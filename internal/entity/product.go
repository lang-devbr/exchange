package entity

import (
	"errors"
	"time"
	"gorm.io/gorm"
)

type Product struct {
	ID         uint         `gorm:"primaryKey" json:"id"`
	Name       string         `json:"name"`
	Price      float64        `json:"price"`
	Created_At time.Time      `json:"created_At"`
	Currencies []Currency     `gorm:"many2many:product_currency;foreignkey:ProductID;" json:"currencies"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func NewProduct(name string, price float64, currencies []Currency) (Product, error) {
	p := Product{
		Name:       name,
		Price:      price,
		Created_At: time.Now(),
		Currencies: currencies,
	}
	err := p.Validate()
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (p Product) Validate() error {
	if p.Name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}
