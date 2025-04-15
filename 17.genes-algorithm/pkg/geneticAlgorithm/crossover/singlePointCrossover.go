package crossover

import (
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
	"math/rand"
)

type SinglePointCrossover struct {
	geneticalgorithm.BaseCrossoverStrategy
}

// 隨機選擇一個交配點（可能是某個索引值），將兩個個體的基因在此點進行交換
func (c *SinglePointCrossover) Crossover(parent1, parent2 geneticalgorithm.Individual) [2]geneticalgorithm.Individual {
	randPivot := rand.Intn(len(parent1.GetChromosome()))

	chromosome1 := parent1.GetChromosome()
	chromosome2 := parent2.GetChromosome()

	offspring1 := append(chromosome1[:randPivot], chromosome2[randPivot:]...)
	offspring2 := append(chromosome2[:randPivot], chromosome1[randPivot:]...)

	return [2]geneticalgorithm.Individual{
		c.BaseCrossoverStrategy.IndividualFactory.CreateIndividual(offspring1),
		c.BaseCrossoverStrategy.IndividualFactory.CreateIndividual(offspring2),
	}
}

func NewSinglePointCrossover(individualFactory geneticalgorithm.IndividualFactory) *SinglePointCrossover {
	c := &SinglePointCrossover{
		BaseCrossoverStrategy: *geneticalgorithm.NewBaseCrossoverStrategy(individualFactory),
	}

	c.BaseCrossoverStrategy.CrossoverStrategy = c
	return c
}

var _ geneticalgorithm.CrossoverStrategy = &SinglePointCrossover{}
