package main

type TopToBottomSortingStrategy struct{}

func (s *TopToBottomSortingStrategy) Sorting(users []Individual) []Individual {
	return users
}

func NewTopToBottomSortingStrategy() *TopToBottomSortingStrategy {
	return &TopToBottomSortingStrategy{}
}
