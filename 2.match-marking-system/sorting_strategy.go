package main

type ISortingStrategy interface {
	Sorting(users []Individual) []Individual
}

var _ ISortingStrategy = &TopToBottomSortingStrategy{}
var _ ISortingStrategy = &BottomToTopSortingStrategy{}
