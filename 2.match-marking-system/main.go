package main

import "fmt"

func main() {
	// 配對策略
	matchingStrategy := NewDistanceMatchingStrategy()

	// 排序策略
	sortingStrategy := NewTopToBottomSortingStrategy()

	// 初始化配對系統
	system := NewMatchmakingSystem(matchingStrategy, sortingStrategy)

	// 建立假對象資料
	users := generateFakeIndividual(10)

	// 為每個用戶找到最佳配對
	for _, user := range users {
		otherUsers := make([]Individual, 0)
		for _, other := range users {
			if user.GetId() != other.GetId() {
				otherUsers = append(otherUsers, other)
			}
		}

		// 進行配對
		result := system.Match(user, otherUsers)
		fmt.Printf("對象 %d (位置: %d,%d) 最佳對象: 用戶 %d (位置: %d,%d)\n",
			user.GetId(),
			user.GetCoord().GetX(),
			user.GetCoord().GetY(),
			result[0].GetId(),
			result[0].GetCoord().GetX(),
			result[0].GetCoord().GetY())
	}
}
