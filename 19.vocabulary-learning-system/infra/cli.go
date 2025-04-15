package infra

import (
	"bufio"
	"fmt"
	"github/elliot9/class19/pkg/sceneManagement"
	"os"
	"strings"
)

type ConsoleIO struct {
	reader *bufio.Reader
}

func NewConsoleIO() *ConsoleIO {
	return &ConsoleIO{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (c *ConsoleIO) ReadLine() string {
	input, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (c *ConsoleIO) Println(s string) {
	fmt.Println(s)
}

type MockCLI struct {
	inputs  []string
	outputs []string
	cursor  int
}

func NewMockCLI(inputs []string) *MockCLI {
	return &MockCLI{
		inputs:  inputs,
		outputs: make([]string, 0),
		cursor:  0,
	}
}

func (m *MockCLI) ReadLine() string {
	if m.cursor >= len(m.inputs) {
		return ""
	}
	input := m.inputs[m.cursor]
	m.cursor++
	return input
}

func (m *MockCLI) Println(s string) {
	m.outputs = append(m.outputs, s)
}

func (m *MockCLI) GetOutputs() []string {
	return m.outputs
}

var _ sceneManagement.CLI = &MockCLI{}
var _ sceneManagement.CLI = &ConsoleIO{}
