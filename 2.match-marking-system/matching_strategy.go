package main

type IMatchingStrategy interface {
	Matching(user Individual, targetUsers []Individual) []Individual
}

var _ IMatchingStrategy = &DistanceMatchingStrategy{}
var _ IMatchingStrategy = &HabitMatchingStrategy{}
