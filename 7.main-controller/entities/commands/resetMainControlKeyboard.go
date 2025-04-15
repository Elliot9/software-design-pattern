package commands

import (
	"github/elliot9/class7/interfaces"
)

type ResetMainControlKeyboardCommand struct {
	controller   interfaces.Controller
	tempKeyboard map[string]interfaces.Command
}

func (c *ResetMainControlKeyboardCommand) Execute() {
	c.tempKeyboard = c.controller.GetKeyboard()
	c.controller.Reset()
}

func (c *ResetMainControlKeyboardCommand) Undo() {
	c.controller.SetKeyboard(c.tempKeyboard)
}

func (c *ResetMainControlKeyboardCommand) GetName() string {
	return "ResetMainControlKeyboard"
}

func NewResetMainControlKeyboardCommand(controller interfaces.Controller) *ResetMainControlKeyboardCommand {
	return &ResetMainControlKeyboardCommand{
		controller: controller,
	}
}

var _ interfaces.Command = &ResetMainControlKeyboardCommand{}
