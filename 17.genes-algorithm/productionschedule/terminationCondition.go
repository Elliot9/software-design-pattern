package productionschedule

import geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"

type TerminationCondition struct{}

// schedule 是個體, 其中包含了很多的 production(基因)
func (t *TerminationCondition) ShouldTerminate(population []geneticalgorithm.Individual) bool {
	// 檢查群體中, 是否有任一個體, 滿足所有需求
	for _, schedule := range population {
		if t.IsSatisfied(schedule) {
			return true
		}
	}
	return false
}

func (t *TerminationCondition) IsSatisfied(schedule geneticalgorithm.Individual) bool {
	// 檢查排程所生產的產品是否滿足所有需求
	for _, requirement := range ProductionScheduleInstance.ProductionRequirements {
		count := 0
		for _, production := range schedule.GetChromosome() {
			production := production.(*Production)
			if production.Product.Name == requirement.Product.Name {
				count += production.Number
			}
		}
		if count < requirement.RequiredNumber {
			return false
		}
	}
	return true
}

var _ geneticalgorithm.TerminationCondition = &TerminationCondition{}

func NewTerminationCondition() *TerminationCondition {
	return &TerminationCondition{}
}
