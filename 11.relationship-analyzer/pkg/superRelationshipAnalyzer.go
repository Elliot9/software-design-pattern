package pkg

import (
	"fmt"
	"strings"
)

type SuperRelationshipAnalyzer struct {
	relationships map[string]map[string]bool
}

func NewSuperRelationshipAnalyzer() *SuperRelationshipAnalyzer {
	return &SuperRelationshipAnalyzer{
		relationships: make(map[string]map[string]bool),
	}
}

func (s *SuperRelationshipAnalyzer) Init(script string) error {
	if script == "" {
		return fmt.Errorf("script is empty")
	}

	lines := strings.Split(script, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " -- ")
		if len(parts) != 2 {
			return fmt.Errorf("invalid line: %s", line)
		}

		name := parts[0]
		friend := parts[1]
		if s.relationships[name] == nil {
			s.relationships[name] = make(map[string]bool)
		}
		if s.relationships[friend] == nil {
			s.relationships[friend] = make(map[string]bool)
		}

		s.relationships[name][friend] = true
		s.relationships[friend][name] = true
	}
	return nil
}

func (s *SuperRelationshipAnalyzer) IsMutualFriend(targetName, name1, name2 string) bool {
	return s.relationships[targetName][name1] && s.relationships[targetName][name2]
}
