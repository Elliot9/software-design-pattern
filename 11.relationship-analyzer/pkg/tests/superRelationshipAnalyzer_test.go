package tests

import (
	"github/elliot9/class11/pkg"
	"testing"
)

func TestSuperRelationshipAnalyzer(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		targetName   string
		name1        string
		name2        string
		expectError  bool
		expectMutual bool
	}{
		{
			name:         "朋友關係",
			input:        "Alice -- Bob\nAlice -- Charlie\nBob -- Charlie",
			targetName:   "Alice",
			name1:        "Bob",
			name2:        "Charlie",
			expectError:  false,
			expectMutual: true,
		},
		{
			name:         "非朋友關係",
			input:        "Alice -- Bob\nCharlie -- David",
			targetName:   "Alice",
			name1:        "Bob",
			name2:        "Charlie",
			expectError:  false,
			expectMutual: false,
		},
		{
			name:         "格式錯誤",
			input:        "Alice - Bob",
			targetName:   "Alice",
			name1:        "Bob",
			name2:        "Charlie",
			expectError:  true,
			expectMutual: false,
		},
		{
			name:         "空字串",
			input:        "",
			expectError:  true,
			expectMutual: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := pkg.NewSuperRelationshipAnalyzer()
			err := analyzer.Init(tt.input)

			// 檢查錯誤情況
			if tt.expectError {
				if err == nil {
					t.Errorf("期望得到錯誤，但是沒有")
				}
				return
			}

			if err != nil {
				t.Errorf("初始化出錯: %v", err)
				return
			}

			// 測試互為朋友的情況
			result := analyzer.IsMutualFriend(tt.targetName, tt.name1, tt.name2)
			if result != tt.expectMutual {
				t.Errorf("IsMutualFriend(%s, %s, %s) = %v; 期望 %v",
					tt.targetName, tt.name1, tt.name2, result, tt.expectMutual)
			}
		})
	}
}
