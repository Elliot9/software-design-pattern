package geneticalgorithm

type Individual interface {
	GetChromosome() []Genes
}

type BaseIndividual struct {
	chromosome []Genes
}

func (i *BaseIndividual) GetChromosome() []Genes {
	return i.chromosome
}

var _ Individual = &BaseIndividual{}

func CreateIndividual(chromosome []Genes) *BaseIndividual {
	return &BaseIndividual{
		chromosome: chromosome,
	}
}
