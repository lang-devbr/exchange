package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	Created_At time.Time `json:"created_At"`
}

func NewProduct(name string, price float64) (Product, error) {
	p := Product{
		ID:         uuid.New().String(),
		Name:       name,
		Price:      price,
		Created_At: time.Now(),
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
