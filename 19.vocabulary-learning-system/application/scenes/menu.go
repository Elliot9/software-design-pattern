package scenes

import (
	"fmt"
)

type Menu interface {
	Render() []string
}

type SimpleMenu struct {
	options []Command
}

func NewSimpleMenu(commands []Command) *SimpleMenu {
	return &SimpleMenu{options: commands}
}

func (m *SimpleMenu) Render() []string {
	var lines []string
	for _, opt := range m.options {
		lines = append(lines, fmt.Sprintf("[%s] %s", opt.GetKey(), opt.GetGoal()))
	}
	return lines
}
