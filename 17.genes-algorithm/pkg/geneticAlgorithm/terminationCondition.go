package geneticalgorithm

type TerminationCondition interface {
	ShouldTerminate(population []Individual) bool
}
