package shoppingRecommand

import geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"

// 推薦清單(個體)
type Recommendation struct {
	Items []*RecommendItem
}

func NewRecommendation() *Recommendation {
	return &Recommendation{
		Items: []*RecommendItem{},
	}
}

func (r *Recommendation) GetChromosome() []geneticalgorithm.Genes {
	chromosome := make([]geneticalgorithm.Genes, len(r.Items))
	for i, item := range r.Items {
		chromosome[i] = NewRecommendItem(item.Product, item.Quantity)
	}
	return chromosome
}

func (r *Recommendation) AddRecommendItem(productId int, quantity int) {
	var product *Product
	for _, p := range ShoppingRecommandInstance.Products {
		if p.ID == productId {
			product = p
			break
		}
	}
	if product != nil {
		// 如果已經在 Items 中，則更新數量
		for _, item := range r.Items {
			if item.Product.ID == productId {
				item.Quantity += quantity
				return
			}
		}
		r.Items = append(r.Items, NewRecommendItem(product, quantity))
	}
}

func (r *Recommendation) GetTotalPrice() int {
	totalPrice := 0
	for _, item := range r.Items {
		totalPrice += item.Product.Price * item.Quantity
	}
	return totalPrice
}

func (r *Recommendation) GetTotalWeight() int {
	totalWeight := 0
	for _, item := range r.Items {
		totalWeight += item.Product.Weight * item.Quantity
	}
	return totalWeight
}

var _ geneticalgorithm.Individual = &Recommendation{}
