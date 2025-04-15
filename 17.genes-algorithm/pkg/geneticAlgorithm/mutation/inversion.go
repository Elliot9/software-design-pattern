package mutation

import (
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
)

type Inversion struct {
	geneticalgorithm.BaseMutationStrategy
}

// 將染色體中的一段連續的基因反轉過來
func (m *Inversion) Mutate(individual geneticalgorithm.Individual) geneticalgorithm.Individual {
	newChromosome := individual.GetChromosome()

	for i := 0; i < len(newChromosome); i++ {
		newChromosome[i] = m.BaseMutationStrategy.GenesFactory.CreateInverseGenes(newChromosome[i])
	}

	return m.BaseMutationStrategy.IndividualFactory.CreateIndividual(newChromosome)
}

func NewInversion(individualFactory geneticalgorithm.IndividualFactory, genesFactory geneticalgorithm.GenesFactory) *Inversion {
	m := &Inversion{
		BaseMutationStrategy: *geneticalgorithm.NewBaseMutationStrategy(individualFactory, genesFactory),
	}

	m.BaseMutationStrategy.MutationStrategy = m
	return m
}

var _ geneticalgorithm.MutationStrategy = &Inversion{}
