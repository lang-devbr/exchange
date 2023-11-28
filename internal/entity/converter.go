package entity

type Converter struct {
	Origin string 
	Price float64 
}

func NewConverter(origin string, price float64) Converter {
	c := Converter{
		Origin: origin,
		Price: price,
	}
	return c
}