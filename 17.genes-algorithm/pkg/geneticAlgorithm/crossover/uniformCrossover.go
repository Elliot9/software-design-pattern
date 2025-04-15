package crossover

import (
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
	"math/rand"
)

// 隨機選擇每一個基因，有一半的機率從第一個個體繼承，有一半的機率從第二個個體繼承。
type UniformCrossover struct {
	geneticalgorithm.BaseCrossoverStrategy
}

func (c *UniformCrossover) Crossover(parent1, parent2 geneticalgorithm.Individual) [2]geneticalgorithm.Individual {
	chromosome1 := parent1.GetChromosome()
	chromosome2 := parent2.GetChromosome()

	maxLength := len(chromosome1)
	if len(chromosome2) > maxLength {
		maxLength = len(chromosome2)
	}

	offspring1 := make([]geneticalgorithm.Genes, maxLength)
	offspring2 := make([]geneticalgorithm.Genes, maxLength)

	for i := 0; i < maxLength; i++ {
		if i < len(chromosome1) && i < len(chromosome2) {
			if rand.Float64() < 0.5 {
				offspring1[i] = chromosome1[i]
				offspring2[i] = chromosome2[i]
			} else {
				offspring1[i] = chromosome2[i]
				offspring2[i] = chromosome1[i]
			}
		} else if i < len(chromosome1) {
			offspring1[i] = chromosome1[i]
			offspring2[i] = chromosome1[i]
		} else {
			offspring1[i] = chromosome2[i]
			offspring2[i] = chromosome2[i]
		}
	}

	return [2]geneticalgorithm.Individual{
		c.BaseCrossoverStrategy.IndividualFactory.CreateIndividual(offspring1),
		c.BaseCrossoverStrategy.IndividualFactory.CreateIndividual(offspring2),
	}
}

func NewUniformCrossover(individualFactory geneticalgorithm.IndividualFactory) *UniformCrossover {
	c := &UniformCrossover{
		BaseCrossoverStrategy: *geneticalgorithm.NewBaseCrossoverStrategy(individualFactory),
	}

	c.BaseCrossoverStrategy.CrossoverStrategy = c
	return c
}
