package productionschedule

import (
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
)

type Schedule struct {
	Productions []*Production
}

func NewSchedule() *Schedule {
	return &Schedule{
		Productions: []*Production{},
	}
}

var _ geneticalgorithm.Individual = &Schedule{}

func (s *Schedule) GetChromosome() []geneticalgorithm.Genes {
	chromosome := make([]geneticalgorithm.Genes, len(s.Productions))
	for i, production := range s.Productions {
		chromosome[i] = production
	}
	return chromosome
}

func (s *Schedule) AppendProduction(productName string, number int, workerID int, machineID int) {
	product := ProductionScheduleInstance.Products[productName]
	if product == nil || workerID >= ProductionScheduleInstance.WorkerCount || machineID >= ProductionScheduleInstance.MachineCount {
		return
	}

	startTime := 0
	if len(s.Productions) > 0 {
		// 對應 workerID 和 machineID 的生產結束時間
		for _, production := range s.Productions {
			if production.WorkerID == workerID || production.MachineID == machineID {
				startTime = max(startTime, production.EndTime)
			}
		}
	}

	production := NewProduction(product, number, workerID, machineID)
	production.SetTime(startTime)
	s.Productions = append(s.Productions, production)
}

func (s *Schedule) GetTotalTime() int {
	latestEndTime := 0
	for _, production := range s.Productions {
		if production.EndTime > latestEndTime {
			latestEndTime = production.EndTime
		}
	}
	return latestEndTime
}
