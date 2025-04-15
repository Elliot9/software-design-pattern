package shoppingRecommand

import (
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
	"math/rand"
)

type IndividualFactory struct{}

func (f *IndividualFactory) CreateIndividual(genes []geneticalgorithm.Genes) geneticalgorithm.Individual {
	recommendation := NewRecommendation()
	for _, gene := range genes {
		recommandItem := gene.(*RecommendItem)
		recommendation.AddRecommendItem(recommandItem.Product.ID, recommandItem.Quantity)
	}
	return recommendation
}

type GenesFactory struct{}

func (f *GenesFactory) CreateRandomGenes() geneticalgorithm.Genes {
	pivot := rand.Intn(len(ShoppingRecommandInstance.Products))
	product := ShoppingRecommandInstance.Products[pivot]

	return NewRecommendItem(product, rand.Intn(2))
}

func (f *GenesFactory) CreateInverseGenes(genes geneticalgorithm.Genes) geneticalgorithm.Genes {
	recommandItem := genes.(*RecommendItem)
	newQuantity := max(1, recommandItem.Quantity+rand.Intn(200)-100)

	// 改變商品 或 數量
	if rand.Intn(2) == 0 {
		for _, product := range ShoppingRecommandInstance.Products {
			if recommandItem.Product.ID == product.ID {
				continue
			}
			return NewRecommendItem(product, newQuantity)
		}
	} else {
		return NewRecommendItem(recommandItem.Product, newQuantity)
	}
	return nil
}

func NewIndividualFactory() *IndividualFactory {
	return &IndividualFactory{}
}

func NewGenesFactory() *GenesFactory {
	return &GenesFactory{}
}

var _ geneticalgorithm.IndividualFactory = &IndividualFactory{}
var _ geneticalgorithm.GenesFactory = &GenesFactory{}
