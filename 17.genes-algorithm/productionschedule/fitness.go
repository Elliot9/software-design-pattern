package productionschedule

import (
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
	"math"
)

type Fitness struct{}

/*
指標：
1. 需求滿足程度
2. 生產時間效率
3. 過量生產懲罰

計算步驟：
1. 統計每種產品的生產數量。
2. 計算需求滿足率：比較生產數量與需求數量。
3. 計算滿足率與需求是否達成。
4. 計算理論最短生產時間：總工作量 = 需求數量 × 生產時間。
5. 計算瓶頸資源 = min(工人數, 機器數)。
6. 計算理論時間 = 總工作量 / 瓶頸資源。
7. 計算時間效率比率：Min(理論最短生產時間 / 總生產時間, 1.0)。
8. 計算超額生產懲罰：若生產超過需求，根據超額比例扣分。
9. 最終配適度計算：若所有需求達成： 0.7 + 0.3 × 時間效率 - 超額懲罰。
10. 限制配適度在 [0,1] 之間。
*/
func (f *Fitness) CalculateFitness(individual geneticalgorithm.Individual) float64 {
	schedule := individual.(*Schedule)

	productionCounts := make(map[string]int)
	for _, production := range schedule.Productions {
		productionCounts[production.Product.Name] += production.Number
	}

	requirementsMet := true
	requirementSatisfactionRatio := 0.0
	totalRequiredProducts := 0
	totalProducedOfRequired := 0

	for _, requirement := range ProductionScheduleInstance.ProductionRequirements {
		productName := requirement.Product.Name
		produced := productionCounts[productName]
		required := requirement.RequiredNumber

		totalRequiredProducts += required
		totalProducedOfRequired += min(produced, required)

		if produced < required {
			requirementsMet = false
		}
	}

	if totalRequiredProducts > 0 {
		requirementSatisfactionRatio = float64(totalProducedOfRequired) / float64(totalRequiredProducts)
	}

	totalTime := schedule.GetTotalTime()
	if totalTime == 0 {
		return 0
	}

	totalWork := 0
	for _, requirement := range ProductionScheduleInstance.ProductionRequirements {
		totalWork += requirement.RequiredNumber * requirement.Product.CostTime
	}

	bottleneckResource := min(ProductionScheduleInstance.WorkerCount, ProductionScheduleInstance.MachineCount)
	theoreticalMinTime := float64(totalWork) / float64(bottleneckResource)

	timeEfficiencyRatio := math.Min(theoreticalMinTime/float64(totalTime), 1.0)

	overproductionPenalty := 0.0
	for _, requirement := range ProductionScheduleInstance.ProductionRequirements {
		productName := requirement.Product.Name
		produced := productionCounts[productName]
		required := requirement.RequiredNumber

		if produced > required {
			excessRatio := float64(produced-required) / float64(required)
			overproductionPenalty += excessRatio * 0.1
		}
	}

	fitness := 0.0

	if requirementsMet {
		fitness = 0.7 + (0.3 * timeEfficiencyRatio) - overproductionPenalty
	} else {
		fitness = requirementSatisfactionRatio * 0.7
	}

	return math.Max(0.0, math.Min(fitness, 1.0))
}

func NewFitness() *Fitness {
	return &Fitness{}
}

var _ geneticalgorithm.FitnessStrategy = &Fitness{}
