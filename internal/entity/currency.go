package entity

import "errors"

type Currency struct {
	Code string `json:"code"`
}

func NewCurrency(code string) (Currency, error) {
	c := Currency{
		Code: code,
	}
	err := c.Validate()
	if err != nil {
		return Currency{}, err
	}
	return c, nil
}

func (c Currency) Validate() error {
	if c.Code == "" {
		return errors.New("code cannot be empty")
	}
	return nil
}