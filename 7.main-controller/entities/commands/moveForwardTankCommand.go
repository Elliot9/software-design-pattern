package commands

import (
	"github/elliot9/class7/entities/devices"
	"github/elliot9/class7/interfaces"
)

type MoveForwardTankCommand struct {
	tank *devices.Tank
}

func (c *MoveForwardTankCommand) Execute() {
	c.tank.MoveForward()
}

func (c *MoveForwardTankCommand) Undo() {
	c.tank.MoveBackward()
}

func (c *MoveForwardTankCommand) GetName() string {
	return "MoveTankForward"
}

func NewMoveForwardTankCommand(tank *devices.Tank) interfaces.Command {
	return &MoveForwardTankCommand{tank: tank}
}

var _ interfaces.Command = &MoveForwardTankCommand{}
