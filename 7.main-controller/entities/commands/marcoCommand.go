package commands

import (
	"github/elliot9/class7/interfaces"
	"strings"
)

type MarcoCommand struct {
	commands []interfaces.Command
}

func (c *MarcoCommand) Execute() {
	for _, command := range c.commands {
		command.Execute()
	}
}

func (c *MarcoCommand) Undo() {
	for i := len(c.commands) - 1; i >= 0; i-- {
		c.commands[i].Undo()
	}
}

func (c *MarcoCommand) GetName() string {
	commandNames := []string{}
	for _, command := range c.commands {
		commandNames = append(commandNames, command.GetName())
	}
	return strings.Join(commandNames, " & ")
}

func (c *MarcoCommand) SetCommands(commands []interfaces.Command) {
	c.commands = commands
}

var _ interfaces.Command = &MarcoCommand{}
