package adapter

import (
	"fmt"
	"github/elliot9/class11/internal/core"
	"github/elliot9/class11/pkg"
	"sort"
	"strings"
)

type RelationshipAnalyzerAdapter struct {
	analyzer *pkg.SuperRelationshipAnalyzer
	users    map[string]bool
}

func NewRelationshipAnalyzerAdapter() *RelationshipAnalyzerAdapter {
	return &RelationshipAnalyzerAdapter{
		analyzer: pkg.NewSuperRelationshipAnalyzer(),
		users:    make(map[string]bool),
	}
}

func (a *RelationshipAnalyzerAdapter) Parse(script string) (core.RelationshipGraph, error) {
	if script == "" {
		return nil, fmt.Errorf("script is empty")
	}

	relationships := []string{}
	connections := [][2]int{}

	lines := strings.Split(script, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		name := parts[0]
		friends := strings.Split(parts[1], " ")

		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		a.users[name] = true

		for _, friend := range friends {
			if friend == "" {
				continue
			}
			a.users[friend] = true
			connections = append(connections, [2]int{ParseUserToIndex(name), ParseUserToIndex(friend)})
			relationships = append(relationships, name+" -- "+friend)
		}
	}

	err := a.analyzer.Init(strings.Join(relationships, "\n"))
	if err != nil {
		return nil, err
	}

	return NewRelationshipGraphAdapter(connections), nil
}

func (a *RelationshipAnalyzerAdapter) GetMutualFriends(name1, name2 string) []string {
	mutualFriends := []string{}
	for user := range a.users {
		if user != name1 && user != name2 && a.analyzer.IsMutualFriend(user, name1, name2) {
			mutualFriends = append(mutualFriends, user)
		}
	}

	sort.Strings(mutualFriends)
	return mutualFriends
}

var _ core.RelationshipAnalyzer = &RelationshipAnalyzerAdapter{}
