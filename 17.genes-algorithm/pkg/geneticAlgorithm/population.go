package geneticalgorithm

type Population interface {
	GetIndividuals() []Individual
}

type BasePopulation struct {
	individuals []Individual
}

func (p *BasePopulation) GetIndividuals() []Individual {
	return p.individuals
}

var _ Population = &BasePopulation{}

func CreatePopulation(individuals []Individual) *BasePopulation {
	return &BasePopulation{
		individuals: individuals,
	}
}
