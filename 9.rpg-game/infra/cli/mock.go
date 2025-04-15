package cli

import (
	"fmt"
	"strconv"
	"strings"
)

type MockCLI struct {
	inputs       []string
	currentIndex int
	outputs      []string
}

func NewMockCLI(inputs []string) *MockCLI {
	return &MockCLI{
		inputs:  inputs,
		outputs: make([]string, 0),
	}
}

func (m *MockCLI) ReadLine() string {
	if m.currentIndex >= len(m.inputs) {
		return ""
	}
	input := m.inputs[m.currentIndex]
	m.currentIndex++
	return input
}

func (m *MockCLI) ReadNumber() []int {
	input := m.ReadLine()
	numbers := []int{}
	for _, number := range strings.Split(input, ", ") {
		num, err := strconv.Atoi(number)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	// fmt.Println("ReadNumber", numbers)
	return numbers
}

func (m *MockCLI) Println(s string) {
	m.outputs = append(m.outputs, s)
	fmt.Println(s)
}

func (m *MockCLI) GetOutputs() []string {
	return m.outputs
}

var _ CLI = &MockCLI{}
