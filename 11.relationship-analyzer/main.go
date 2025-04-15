package main

import (
	"fmt"
	"github/elliot9/class11/internal/adapter"
)

func main() {
	relationshipAnalyzer := adapter.NewRelationshipAnalyzerAdapter()

	script := `A: B C D
B: A D E
C: A E
F: G`
	graph, _ := relationshipAnalyzer.Parse(script)

	// 共同好友
	mutualFriends := relationshipAnalyzer.GetMutualFriends("A", "B")
	fmt.Println(mutualFriends)

	// 是否連結
	connected := graph.HasConnection("A", "F")
	fmt.Println(connected)
}
