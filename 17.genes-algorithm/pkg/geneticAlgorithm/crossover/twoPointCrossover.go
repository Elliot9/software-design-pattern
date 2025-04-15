package crossover

import (
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
	"math/rand"
)

type TwoPointCrossover struct {
	geneticalgorithm.BaseCrossoverStrategy
}

// 隨機選擇兩個交配點，將兩個個體的基因在這兩點之間的區間進行交換。
func (c *TwoPointCrossover) Crossover(parent1, parent2 geneticalgorithm.Individual) [2]geneticalgorithm.Individual {
	chromosomeLength := len(parent1.GetChromosome())
	if chromosomeLength < 2 {
		// 如果染色體長度小於 2，則無法進行雙點交配
		return [2]geneticalgorithm.Individual{parent1, parent2}
	}

	randPivot1, randPivot2 := rand.Intn(chromosomeLength-1), rand.Intn(chromosomeLength-1)
	for randPivot1 == randPivot2 {
		randPivot2 = rand.Intn(chromosomeLength - 1)
	}

	if randPivot1 > randPivot2 {
		randPivot1, randPivot2 = randPivot2, randPivot1
	}

	chromosome1 := parent1.GetChromosome()
	chromosome2 := parent2.GetChromosome()

	offspring1 := append(append(chromosome1[:randPivot1], chromosome2[randPivot1:randPivot2]...), chromosome1[randPivot2:]...)
	offspring2 := append(append(chromosome2[:randPivot1], chromosome1[randPivot1:randPivot2]...), chromosome2[randPivot2:]...)

	return [2]geneticalgorithm.Individual{
		c.BaseCrossoverStrategy.IndividualFactory.CreateIndividual(offspring1),
		c.BaseCrossoverStrategy.IndividualFactory.CreateIndividual(offspring2),
	}
}

func NewTwoPointCrossover(individualFactory geneticalgorithm.IndividualFactory) *TwoPointCrossover {
	c := &TwoPointCrossover{
		BaseCrossoverStrategy: *geneticalgorithm.NewBaseCrossoverStrategy(individualFactory),
	}

	c.BaseCrossoverStrategy.CrossoverStrategy = c
	return c
}

var _ geneticalgorithm.CrossoverStrategy = &TwoPointCrossover{}
