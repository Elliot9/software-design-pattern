package shoppingRecommand

import geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"

// 清單項目(基因)
type RecommendItem struct {
	Product  *Product
	Quantity int
}

func NewRecommendItem(product *Product, quantity int) *RecommendItem {
	return &RecommendItem{
		Product:  product,
		Quantity: quantity,
	}
}

func (g *RecommendItem) GetValue() any {
	return g
}

var _ geneticalgorithm.Genes = &RecommendItem{}
