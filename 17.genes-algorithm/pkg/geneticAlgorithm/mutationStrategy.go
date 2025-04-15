package geneticalgorithm

type MutationStrategy interface {
	Mutate(individual Individual) Individual
}

type BaseMutationStrategy struct {
	MutationStrategy
	IndividualFactory IndividualFactory
	GenesFactory      GenesFactory
}

func NewBaseMutationStrategy(individualFactory IndividualFactory, genesFactory GenesFactory) *BaseMutationStrategy {
	return &BaseMutationStrategy{
		IndividualFactory: individualFactory,
		GenesFactory:      genesFactory,
	}
}

var _ MutationStrategy = &BaseMutationStrategy{}
