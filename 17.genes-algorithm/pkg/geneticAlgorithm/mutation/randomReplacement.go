package mutation

import (
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
	"math/rand"
)

type RandomReplacement struct {
	geneticalgorithm.BaseMutationStrategy
}

// 是指對染色體中的一個或多個基因「隨機地」替換成新的基因值。
func (m *RandomReplacement) Mutate(individual geneticalgorithm.Individual) geneticalgorithm.Individual {
	newChromosome := individual.GetChromosome()
	mutationCount := rand.Intn(len(newChromosome))

	for i := 0; i < mutationCount; i++ {
		randPivot := rand.Intn(len(newChromosome))
		newChromosome[randPivot] = m.BaseMutationStrategy.GenesFactory.CreateRandomGenes()
	}

	return m.BaseMutationStrategy.IndividualFactory.CreateIndividual(newChromosome)
}

func NewRandomReplacement(individualFactory geneticalgorithm.IndividualFactory, genesFactory geneticalgorithm.GenesFactory) *RandomReplacement {
	m := &RandomReplacement{
		BaseMutationStrategy: *geneticalgorithm.NewBaseMutationStrategy(individualFactory, genesFactory),
	}

	m.BaseMutationStrategy.MutationStrategy = m
	return m
}
