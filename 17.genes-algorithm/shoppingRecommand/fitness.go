package shoppingRecommand

import (
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
	"math"
)

type Fitness struct{}

/*
指標：
1. 商品符合顧客偏好的程度。
2. 預算與重量的利用率。
3. 符合顧客的預算與負重限制

計算步驟：
 1. 計算推薦方案的總價格與總重量。
 2. 若總價格超過顧客預算或總重量超過可負重，則配適度為 0.0。
 3. 計算商品偏好得分：
 4. 依據顧客對每個產品類別的偏好值進行加權計算。
 5. 計算預算與重量的利用率：
 6. 預算利用率 = 總價格 / 顧客預算。
 7. 負重利用率 = 總重量 / 顧客可負重。
 8. 綜合評估配適度：
    60% 商品偏好得分。
    20% 預算利用率。
    20% 負重利用率。
 9. 限制配適度在 [0,1] 之間
*/
func (f *Fitness) CalculateFitness(individual geneticalgorithm.Individual) float64 {
	recommendation := individual.(*Recommendation)

	totalPrice := recommendation.GetTotalPrice()
	totalWeight := recommendation.GetTotalWeight()

	if totalPrice > ShoppingRecommandInstance.Customer.Budget ||
		totalWeight > ShoppingRecommandInstance.Customer.Capacity {
		return 0.0
	}

	preferenceScore := 0.0
	totalItems := 0

	for _, item := range recommendation.Items {
		preference := ShoppingRecommandInstance.Customer.Preferences[item.Product.Category]
		preferenceScore += preference * float64(item.Quantity)
		totalItems += item.Quantity
	}

	if totalItems == 0 {
		return 0.0
	}

	avgPreferenceScore := preferenceScore / float64(totalItems)

	budgetUtilization := float64(totalPrice) / float64(ShoppingRecommandInstance.Customer.Budget)
	capacityUtilization := float64(totalWeight) / float64(ShoppingRecommandInstance.Customer.Capacity)

	fitness := 0.6*avgPreferenceScore +
		0.2*budgetUtilization +
		0.2*capacityUtilization

	return math.Min(1.0, fitness)
}

func NewFitness() *Fitness {
	return &Fitness{}
}

var _ geneticalgorithm.FitnessStrategy = &Fitness{}
