package adapter

import (
	"github/elliot9/class11/internal/core"

	"github.com/pengubco/algorithms/union_find"
)

const (
	MaxUsers = 1000
)

type RelationshipGraphAdapter struct {
	graph *union_find.UnionFind
}

func NewRelationshipGraphAdapter(relationships [][2]int) *RelationshipGraphAdapter {
	graph := union_find.NewUnionFind(MaxUsers)
	for _, relationship := range relationships {
		graph.Union(relationship[0], relationship[1])
	}

	return &RelationshipGraphAdapter{graph: graph}
}

func (a *RelationshipGraphAdapter) HasConnection(name1, name2 string) bool {
	return a.graph.Find(ParseUserToIndex(name1)) == a.graph.Find(ParseUserToIndex(name2))
}

func ParseUserToIndex(name string) int {
	return int(name[0] - '0')
}

var _ core.RelationshipGraph = &RelationshipGraphAdapter{}
