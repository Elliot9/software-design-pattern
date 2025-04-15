package geneticalgorithm

type IndividualFactory interface {
	CreateIndividual(genes []Genes) Individual
}

type GenesFactory interface {
	CreateRandomGenes() Genes
	CreateInverseGenes(genes Genes) Genes
}
