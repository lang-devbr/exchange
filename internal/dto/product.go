package dto

import (
	"time"

	"github.com/lang-dev/exchange/internal/entity"
	"gorm.io/gorm"
)

type ProductRequest struct {
	ID         uint              `gorm:"primaryKey" json:"id"`
	Name       string            `json:"name"`
	Price      float64           `json:"price"`
	Currencies []entity.Currency `gorm:"many2many:product_currency;" json:"currencies"`
	CreatedAt time.Time	
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
