package main

type MatchmakingSystem struct {
	matchingStrategy IMatchingStrategy
	sortingStrategy  ISortingStrategy
}

func (m *MatchmakingSystem) Match(user Individual, targetUsers []Individual) []Individual {
	matchingUsers := m.matchingStrategy.Matching(user, targetUsers)
	sortingUsers := m.sortingStrategy.Sorting(matchingUsers)
	return sortingUsers
}

func NewMatchmakingSystem(matchingStrategy IMatchingStrategy, sortingStrategy ISortingStrategy) *MatchmakingSystem {
	return &MatchmakingSystem{matchingStrategy: matchingStrategy, sortingStrategy: sortingStrategy}
}
