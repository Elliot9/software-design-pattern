package infra

type CLI interface {
	Println(str string)
	GetOutput() []string
}

type MockCLI struct {
	currentIndex int
	Output       []string
}

func NewMockCLI() *MockCLI {
	return &MockCLI{Output: make([]string, 0)}
}

func (m *MockCLI) Println(str string) {
	m.Output = append(m.Output, str)
	m.currentIndex++
}

func (m *MockCLI) GetOutput() []string {
	return m.Output
}
