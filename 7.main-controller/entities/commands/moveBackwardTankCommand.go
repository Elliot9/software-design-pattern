package commands

import (
	"github/elliot9/class7/entities/devices"
	"github/elliot9/class7/interfaces"
)

type MoveBackwardTankCommand struct {
	tank *devices.Tank
}

func (c *MoveBackwardTankCommand) Execute() {
	c.tank.MoveBackward()
}

func (c *MoveBackwardTankCommand) Undo() {
	c.tank.MoveForward()
}

func (c *MoveBackwardTankCommand) GetName() string {
	return "MoveTankBackward"
}

func NewMoveBackwardTankCommand(tank *devices.Tank) interfaces.Command {
	return &MoveBackwardTankCommand{tank: tank}
}

var _ interfaces.Command = &MoveBackwardTankCommand{}
