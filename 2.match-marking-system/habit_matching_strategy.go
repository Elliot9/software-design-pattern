package main

import (
	"sort"
)

type HabitMatchingStrategy struct{}

func (h *HabitMatchingStrategy) Matching(user Individual, targetUsers []Individual) []Individual {
	var sameHabitCount = func(userA, userB Individual) int {
		count := 0
		for _, habit := range userA.GetHabits() {
			for _, habit2 := range userB.GetHabits() {
				if habit.GetName() == habit2.GetName() {
					count++
				}
			}
		}
		return count
	}

	sort.Slice(targetUsers, func(i, j int) bool {
		user1 := targetUsers[i]
		user2 := targetUsers[j]
		count1 := sameHabitCount(user, user1)
		count2 := sameHabitCount(user, user2)
		if count1 == count2 {
			return user1.GetId() < user2.GetId()
		}
		return count1 > count2
	})
	return targetUsers
}

func NewHabitMatchingStrategy() *HabitMatchingStrategy {
	return &HabitMatchingStrategy{}
}
