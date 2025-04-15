package geneticalgorithm

type GeneticAlgorithm struct {
	MaxNumberOfIterations int
	TerminationCondition  TerminationCondition
	Selection             SelectionStrategy
	Crossover             CrossoverStrategy
	Mutation              MutationStrategy
	Fitness               FitnessStrategy
}

func (ga *GeneticAlgorithm) Optimize(population Population) Individual {
	currentPopulation := population.GetIndividuals()

	for i := 0; i < ga.MaxNumberOfIterations; i++ {
		parent1, parent2 := ga.Selection.Select(currentPopulation)
		offsprings := ga.Crossover.Crossover(parent1, parent2)
		newPopulation := make([]Individual, 0)
		for _, offspring := range offsprings {
			newPopulation = append(newPopulation, ga.Mutation.Mutate(offspring))
		}

		currentPopulation = append(currentPopulation, newPopulation...)

		if ga.TerminationCondition != nil && ga.TerminationCondition.ShouldTerminate(currentPopulation) {
			return ga.findBestIndividual(currentPopulation)
		}
	}

	return ga.findBestIndividual(currentPopulation)
}

func (ga *GeneticAlgorithm) findBestIndividual(population []Individual) Individual {
	bestIndividual := population[0]
	for _, individual := range population {
		if ga.Fitness.CalculateFitness(individual) > ga.Fitness.CalculateFitness(bestIndividual) {
			bestIndividual = individual
		}
	}
	return bestIndividual
}
