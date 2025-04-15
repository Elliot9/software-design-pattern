package factory

import "github/elliot9/class10/internal/core"

type OutPutPrescriptionStrategyFactory struct {
	strategies map[string]func() core.OutPutPrescriptionStrategy
}

func NewOutPutPrescriptionStrategyFactory() *OutPutPrescriptionStrategyFactory {
	return &OutPutPrescriptionStrategyFactory{
		strategies: make(map[string]func() core.OutPutPrescriptionStrategy),
	}
}

func (f *OutPutPrescriptionStrategyFactory) Register(name string, creator func() core.OutPutPrescriptionStrategy) {
	f.strategies[name] = creator
}

func (f *OutPutPrescriptionStrategyFactory) Create(name string) core.OutPutPrescriptionStrategy {
	if creator, exists := f.strategies[name]; exists {
		return creator()
	}
	return nil
}
