package geneticalgorithm

type SelectionStrategy interface {
	Select(population []Individual) (Individual, Individual)
}
