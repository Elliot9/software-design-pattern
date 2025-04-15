package core

type RelationshipAnalyzer interface {
	Parse(script string) (RelationshipGraph, error)
	GetMutualFriends(name1 string, name2 string) []string
}
