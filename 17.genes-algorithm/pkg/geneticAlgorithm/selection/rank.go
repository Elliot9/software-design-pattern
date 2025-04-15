package selection

import (
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
	"math/rand"
	"sort"
)

type Rank struct {
	Fitness geneticalgorithm.FitnessStrategy
}

// 依照適應度大小排序後，編號分別為 3、1、2、4，接著根據個體編號進行抽籤
// 抽籤結果為 3、1，接著將這兩個個體進行交配
func (s *Rank) Select(population []geneticalgorithm.Individual) (geneticalgorithm.Individual, geneticalgorithm.Individual) {
	ranked := make([]geneticalgorithm.Individual, len(population))
	for i, ind := range population {
		ranked[i] = ind
	}

	// 依照適度大小排序
	sort.Slice(ranked, func(i, j int) bool {
		return s.Fitness.CalculateFitness(ranked[i]) < s.Fitness.CalculateFitness(ranked[j])
	})

	// 隨選 2 個個體
	idx1 := rand.Intn(len(ranked))
	idx2 := rand.Intn(len(ranked))
	for idx1 == idx2 {
		idx2 = rand.Intn(len(ranked))
	}

	return ranked[idx1], ranked[idx2]
}

func NewRank(fitness geneticalgorithm.FitnessStrategy) *Rank {
	return &Rank{
		Fitness: fitness,
	}
}

var _ geneticalgorithm.SelectionStrategy = &Rank{}
