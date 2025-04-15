package productionschedule

import (
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
	"math/rand"
)

type IndividualFactory struct{}

func (f *IndividualFactory) CreateIndividual(genes []geneticalgorithm.Genes) geneticalgorithm.Individual {
	schedule := NewSchedule()
	for _, gene := range genes {
		production := gene.(*Production)
		schedule.AppendProduction(production.Product.Name, production.Number, production.WorkerID, production.MachineID)
	}
	return schedule
}

type GenesFactory struct{}

func (f *GenesFactory) CreateRandomGenes() geneticalgorithm.Genes {
	// 隨機選擇一個產品 from Products
	currentPivot, pivot := 0, rand.Intn(len(ProductionScheduleInstance.Products))

	for _, product := range ProductionScheduleInstance.Products {
		if currentPivot == pivot {
			return NewProduction(product, rand.Intn(100), rand.Intn(ProductionScheduleInstance.WorkerCount), rand.Intn(ProductionScheduleInstance.MachineCount))
		}
		currentPivot++
	}

	return nil
}

func (f *GenesFactory) CreateInverseGenes(genes geneticalgorithm.Genes) geneticalgorithm.Genes {
	production := genes.(*Production)
	newNumber := max(1, production.Number+rand.Intn(200)-100)

	// 改變產品對象 或 工作機台
	if rand.Intn(2) == 0 {
		for _, product := range ProductionScheduleInstance.Products {
			if production.Product.Name == product.Name {
				continue
			}
			return NewProduction(product, newNumber, production.WorkerID, production.MachineID)
		}
	} else {
		workerID := rand.Intn(ProductionScheduleInstance.WorkerCount)
		machineID := rand.Intn(ProductionScheduleInstance.MachineCount)
		return NewProduction(production.Product, newNumber, workerID, machineID)
	}

	return nil
}

func NewIndividualFactory() *IndividualFactory {
	return &IndividualFactory{}
}

func NewGenesFactory() *GenesFactory {
	return &GenesFactory{}
}

var _ geneticalgorithm.IndividualFactory = &IndividualFactory{}
var _ geneticalgorithm.GenesFactory = &GenesFactory{}
