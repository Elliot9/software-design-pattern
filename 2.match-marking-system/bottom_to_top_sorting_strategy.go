package main

type BottomToTopSortingStrategy struct{}

// Reverse 是一種 bottom to top 的排序策略
func (r *BottomToTopSortingStrategy) Sorting(users []Individual) []Individual {
	for i, j := 0, len(users)-1; i < j; i, j = i+1, j-1 {
		users[i], users[j] = users[j], users[i]
	}
	return users
}

func NewBottomToTopSortingStrategy() *BottomToTopSortingStrategy {
	return &BottomToTopSortingStrategy{}
}
