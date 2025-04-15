package core

import (
	"bufio"
	"fmt"
	"github/elliot9/class7/entities/commands"
	"github/elliot9/class7/interfaces"
	"os"
	"strconv"
	"strings"
)

type MainController struct {
	commands []interfaces.Command
	keyboard map[string]interfaces.Command
	history  []interfaces.Command
	undo     []interfaces.Command
}

func (c *MainController) Bind(key string, command interfaces.Command) error {
	if len(key) != 1 || key[0] < 'a' || key[0] > 'z' {
		return fmt.Errorf("invalid key")
	}

	c.keyboard[key] = command
	return nil
}

func (c *MainController) Invoke(key string) error {
	command, ok := c.keyboard[key]
	if !ok {
		return fmt.Errorf("command not found")
	}

	c.history = append(c.history, command)
	c.resetUndo()
	command.Execute()
	return nil
}

func (c *MainController) Undo() {
	if len(c.history) == 0 {
		return
	}

	command := c.history[len(c.history)-1]
	command.Undo()

	c.undo = append(c.undo, command)
	c.history = c.history[:len(c.history)-1]
}

func (c *MainController) Redo() {
	if len(c.undo) == 0 {
		return
	}

	command := c.undo[len(c.undo)-1]
	command.Execute()

	c.history = append(c.history, command)
	c.undo = c.undo[:len(c.undo)-1]
}

func (c *MainController) RecordMarco() interfaces.Command {
	inputs := bufio.NewScanner(os.Stdin)
	inputs.Scan()
	keys := strings.Split(inputs.Text(), " ")
	macroCommands := []interfaces.Command{}

	for _, key := range keys {
		commandIndex, err := strconv.Atoi(key)
		if err != nil {
			fmt.Printf("command index %s invalid\n", key)
			continue
		}

		if commandIndex < 0 || commandIndex >= len(c.commands) {
			fmt.Printf("command index %d invalid\n", commandIndex)
			continue
		}

		macroCommands = append(macroCommands, c.commands[commandIndex])
	}

	marcoCommand := &commands.MarcoCommand{}
	marcoCommand.SetCommands(macroCommands)
	return marcoCommand
}

func (c *MainController) resetKeyboard() {
	c.keyboard = make(map[string]interfaces.Command)
}

func (c *MainController) resetUndo() {
	c.undo = []interfaces.Command{}
}

func (c *MainController) GetKeyboard() map[string]interfaces.Command {
	return c.keyboard
}

func (c *MainController) SetKeyboard(keyboard map[string]interfaces.Command) {
	c.keyboard = keyboard
}

func (c *MainController) Reset() {
	c.resetKeyboard()
}

func (c *MainController) PrintfKeyboard() {
	for key, command := range c.keyboard {
		fmt.Printf("%s: %s\n", key, command.GetName())
	}
}

func (c *MainController) PrintfAllCommands() {
	for i, command := range c.commands {
		fmt.Printf("(%d) %s\n", i, command.GetName())
	}
}

func (c *MainController) GetCommands(i int) interfaces.Command {
	return c.commands[i]
}

func (c *MainController) SetCommands(commands []interfaces.Command) {
	c.commands = commands
}

func NewMainController() *MainController {
	return &MainController{
		keyboard: make(map[string]interfaces.Command),
		history:  []interfaces.Command{},
		undo:     []interfaces.Command{},
	}
}

var _ interfaces.Controller = &MainController{}
