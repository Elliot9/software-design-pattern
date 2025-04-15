package main

import (
	"sort"
)

type DistanceMatchingStrategy struct{}

func (d *DistanceMatchingStrategy) Matching(user Individual, targetUsers []Individual) []Individual {
	var twoPointDistance = func(x1, y1, x2, y2 int) int {
		return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	}

	sort.Slice(targetUsers, func(i, j int) bool {
		user1 := targetUsers[i]
		user2 := targetUsers[j]
		distance1 := twoPointDistance(user.GetCoord().GetX(), user.GetCoord().GetY(), user1.GetCoord().GetX(), user1.GetCoord().GetY())
		distance2 := twoPointDistance(user.GetCoord().GetX(), user.GetCoord().GetY(), user2.GetCoord().GetX(), user2.GetCoord().GetY())

		if distance1 == distance2 {
			return user1.GetId() < user2.GetId()
		}

		return distance1 < distance2
	})
	return targetUsers
}

func NewDistanceMatchingStrategy() *DistanceMatchingStrategy {
	return &DistanceMatchingStrategy{}
}
