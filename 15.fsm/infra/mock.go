package infra

type MockCLI struct {
	outputs []string
}

func NewMockCLI() *MockCLI {
	return &MockCLI{
		outputs: make([]string, 0),
	}
}

func (m *MockCLI) Println(s string) {
	m.outputs = append(m.outputs, s)
}

func (m *MockCLI) GetOutputs() []string {
	return m.outputs
}
