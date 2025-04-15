package unit

import (
	"github/elliot9/class11/internal/adapter"
	"reflect"
	"testing"
)

func TestRelationshipAnalyzerAdapter_Parse(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name: "正常輸入測試",
			input: `A: B C D
B: A D E
C: A E`,
			expectError: false,
		},
		{
			name:        "空字串測試",
			input:       "",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			analyzer := adapter.NewRelationshipAnalyzerAdapter()

			_, err := analyzer.Parse(tc.input)

			if tc.expectError {
				if err == nil {
					t.Errorf("期望得到錯誤，但是沒有")
				}
				return
			}

			if err != nil {
				t.Errorf("初始化出錯: %v", err)
				return
			}
		})
	}
}

func TestRelationshipAnalyzerAdapter_GetMutualFriends(t *testing.T) {
	testCases := []struct {
		name            string
		setupScript     string
		person1         string
		person2         string
		expectedFriends []string
	}{
		{
			name: "有共同朋友的情況",
			setupScript: `A: B C D
B: A D E
C: A E
D: A B E
E: B C D`,
			person1:         "B",
			person2:         "C",
			expectedFriends: []string{"A", "E"},
		},
		{
			name: "沒有共同朋友的情況",
			setupScript: `A: B
B: A
C: D
D: C`,
			person1:         "A",
			person2:         "C",
			expectedFriends: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 建立測試對象
			analyzer := adapter.NewRelationshipAnalyzerAdapter()

			_, err := analyzer.Parse(tc.setupScript)
			if err != nil {
				t.Errorf("初始化出錯: %v", err)
				return
			}

			// 執行測試
			result := analyzer.GetMutualFriends(tc.person1, tc.person2)

			if !reflect.DeepEqual(result, tc.expectedFriends) {
				t.Errorf("預期結果: %v, 實際結果: %v", tc.expectedFriends, result)
			}
		})
	}
}
