package selection

import (
	geneticalgorithm "github/elliot9/class17/pkg/geneticAlgorithm"
	"math/rand"
)

type Tournament struct {
	Size    int
	Fitness geneticalgorithm.FitnessStrategy
}

func (s *Tournament) Select(population []geneticalgorithm.Individual) (geneticalgorithm.Individual, geneticalgorithm.Individual) {
	if len(population) < s.Size || s.Size < 4 || s.Size%2 != 0 {
		panic("Tournament size must be an even number and at least 4")
	}

	selected := make([]geneticalgorithm.Individual, s.Size)
	usedIndices := make(map[int]bool)

	for i := 0; i < s.Size; i++ {
		var idx int
		for {
			idx = rand.Intn(len(population))
			if !usedIndices[idx] {
				usedIndices[idx] = true
				break
			}
		}
		selected[i] = population[idx]
	}

	for len(selected) > 2 {
		selected = s.tournamentRound(selected)
	}

	return selected[0], selected[1]
}

// 進行一輪錦標賽，淘汰一半的個體
func (s *Tournament) tournamentRound(group []geneticalgorithm.Individual) []geneticalgorithm.Individual {
	size := len(group)
	winners := make([]geneticalgorithm.Individual, (size+1)/2)

	index := 0
	for i := 0; i < size-1; i += 2 {
		winners[index] = s.tournamentBattle(group[i], group[i+1])
		index++
	}

	if size%2 == 1 {
		winners[index] = group[size-1]
	}

	return winners
}

func (s *Tournament) tournamentBattle(ind1, ind2 geneticalgorithm.Individual) geneticalgorithm.Individual {
	if s.Fitness.CalculateFitness(ind1) > s.Fitness.CalculateFitness(ind2) {
		return ind1
	}
	return ind2
}

func NewTournament(fitness geneticalgorithm.FitnessStrategy, size int) *Tournament {
	return &Tournament{
		Fitness: fitness,
		Size:    size,
	}
}

var _ geneticalgorithm.SelectionStrategy = &Tournament{}
