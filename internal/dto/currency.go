package dto

type CurrencyRequest struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `json:"code"`
}
