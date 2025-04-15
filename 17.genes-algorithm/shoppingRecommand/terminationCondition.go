package shoppingRecommand

import (
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
)

type TerminationCondition struct{}

func (t *TerminationCondition) ShouldTerminate(population []geneticalgorithm.Individual) bool {
	// 檢查群體中, 是否有任一個體, 超過需求限制
	for _, recommand := range population {
		if t.IsOverLimit(recommand) {
			return true
		}
	}
	return false
}

func (t *TerminationCondition) IsOverLimit(recommand geneticalgorithm.Individual) bool {
	recommendation := recommand.(*Recommendation)
	totalPrice := 0
	totalWeight := 0
	for _, item := range recommendation.Items {
		totalPrice += item.Product.Price * item.Quantity
		totalWeight += item.Product.Weight * item.Quantity
	}

	return totalPrice > ShoppingRecommandInstance.Customer.Budget ||
		totalWeight > ShoppingRecommandInstance.Customer.Capacity
}

var _ geneticalgorithm.TerminationCondition = &TerminationCondition{}

func NewTerminationCondition() *TerminationCondition {
	return &TerminationCondition{}
}
