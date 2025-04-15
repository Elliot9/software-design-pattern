package geneticalgorithm

type CrossoverStrategy interface {
	Crossover(parent1, parent2 Individual) [2]Individual
}

type BaseCrossoverStrategy struct {
	IndividualFactory IndividualFactory
	CrossoverStrategy
}

func NewBaseCrossoverStrategy(individualFactory IndividualFactory) *BaseCrossoverStrategy {
	return &BaseCrossoverStrategy{
		IndividualFactory: individualFactory,
	}
}

var _ CrossoverStrategy = &BaseCrossoverStrategy{}
