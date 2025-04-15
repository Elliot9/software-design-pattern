package shoppingRecommand

type Product struct {
	ID       int
	Price    int
	Weight   int
	Category string
}

func NewProduct(id int, price int, weight int, category string) *Product {
	return &Product{
		ID:       id,
		Price:    price,
		Weight:   weight,
		Category: category,
	}
}
