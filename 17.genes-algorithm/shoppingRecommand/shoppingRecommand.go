package shoppingRecommand

type ShoppingRecommand struct {
	Customer *Customer
	Products []*Product
}

var ShoppingRecommandInstance *ShoppingRecommand

func init() {
	ShoppingRecommandInstance = NewShoppingRecommand()
}

func NewShoppingRecommand() *ShoppingRecommand {
	return &ShoppingRecommand{
		Customer: NewCustomer(make(map[string]float64), 0, 0),
		Products: []*Product{},
	}
}

func (s *ShoppingRecommand) AddProduct(price int, weight int, category string) {
	lastID := len(s.Products)
	product := NewProduct(lastID+1, price, weight, category)
	s.Products = append(s.Products, product)
}

func (s *ShoppingRecommand) SetCustomerBudget(budget int) {
	s.Customer.SetBudget(budget)
}

func (s *ShoppingRecommand) SetCustomerCapacity(capacity int) {
	s.Customer.SetCapacity(capacity)
}

func (s *ShoppingRecommand) SetCustomerPreference(category string, preference float64) {
	s.Customer.SetPreference(category, preference)
}
