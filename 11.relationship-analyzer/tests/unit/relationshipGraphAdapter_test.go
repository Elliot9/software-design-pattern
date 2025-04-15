package unit

import (
	"github/elliot9/class11/internal/adapter"
	"testing"
)

func TestNewRelationshipGraphAdapter(t *testing.T) {
	adapter := adapter.NewRelationshipGraphAdapter([][2]int{{0, 1}, {1, 2}, {2, 3}})
	if adapter == nil {
		t.Error("adapter is nil")
	}

	if !adapter.HasConnection("0", "1") {
		t.Error("0 and 1 should be connected")
	}

	if !adapter.HasConnection("1", "2") {
		t.Error("1 and 2 should be connected")
	}

	if !adapter.HasConnection("2", "3") {
		t.Error("2 and 3 should be connected")
	}

	if !adapter.HasConnection("0", "3") {
		t.Error("0 and 3 should be connected")
	}
}

func TestParseUserToIndex(t *testing.T) {
	if adapter.ParseUserToIndex("0") != 0 {
		t.Error("0 should be parsed to 0")
	}
}

func TestNotConnectioned(t *testing.T) {
	adapter := adapter.NewRelationshipGraphAdapter([][2]int{{0, 1}, {1, 4}, {2, 3}})
	if adapter.HasConnection("1", "2") {
		t.Error("1 and 2 should not be connected")
	}

	if adapter.HasConnection("0", "3") {
		t.Error("0 and 3 should not be connected")
	}
}
