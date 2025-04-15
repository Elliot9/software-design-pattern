package integration

import (
	"testing"

	"github/elliot9/class11/internal/adapter"
)

func TestIntegrationScenario(t *testing.T) {
	testCases := []struct {
		name         string
		script       string
		connectioned [][2]string
		disconnected [][2]string
	}{
		{
			name: "完整連結測試",
			script: `A: B C D
B: A D E
C: A E
F: G`,
			connectioned: [][2]string{
				{"A", "D"},
				{"A", "E"},
				{"B", "D"},
				{"B", "E"},
				{"C", "E"},
				{"C", "D"},
			},
			disconnected: [][2]string{
				{"A", "F"},
				{"A", "G"},
				{"E", "F"},
				{"E", "G"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			analyzer := adapter.NewRelationshipAnalyzerAdapter()
			graph, err := analyzer.Parse(tc.script)
			if err != nil {
				t.Fatalf("解析腳本失敗: %v", err)
			}

			// 測試應該有連接的關係
			for _, conn := range tc.connectioned {
				if !graph.HasConnection(conn[0], conn[1]) {
					t.Errorf("期望 %s 和 %s 有連接，但實際上沒有", conn[0], conn[1])
				}
			}

			// 測試應該沒有連接的關係
			for _, conn := range tc.disconnected {
				if graph.HasConnection(conn[0], conn[1]) {
					t.Errorf("期望 %s 和 %s 沒有連接，但實際上有連接", conn[0], conn[1])
				}
			}
		})
	}
}
