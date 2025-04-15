package shoppingRecommand

type Customer struct {
	Preferences map[string]float64
	Budget      int
	Capacity    int
}

func NewCustomer(preferences map[string]float64, budget int, capacity int) *Customer {
	return &Customer{
		Preferences: preferences,
		Budget:      budget,
		Capacity:    capacity,
	}
}

func (c *Customer) SetBudget(budget int) {
	c.Budget = budget
}

func (c *Customer) SetCapacity(capacity int) {
	c.Capacity = capacity
}

func (c *Customer) SetPreference(category string, preference float64) {
	c.Preferences[category] = preference
}
