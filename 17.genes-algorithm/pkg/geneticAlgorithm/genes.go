package geneticalgorithm

type Genes interface {
	GetValue() any
}

type BaseGenes[T any] struct {
	value T
}

func (g *BaseGenes[T]) GetValue() T {
	return g.value
}

var _ Genes = &BaseGenes[any]{}
