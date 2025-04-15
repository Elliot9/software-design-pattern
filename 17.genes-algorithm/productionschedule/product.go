package productionschedule

type Product struct {
	Name     string
	CostTime int
}

func NewProduct(name string, costTime int) *Product {
	return &Product{
		Name:     name,
		CostTime: costTime,
	}
}
