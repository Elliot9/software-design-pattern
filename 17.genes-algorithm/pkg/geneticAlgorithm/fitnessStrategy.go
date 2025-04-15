package geneticalgorithm

type FitnessStrategy interface {
	CalculateFitness(individual Individual) float64
}
