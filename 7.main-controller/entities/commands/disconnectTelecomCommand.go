package commands

import (
	"github/elliot9/class7/entities/devices"
	"github/elliot9/class7/interfaces"
)

type DisconnectTelecomCommand struct {
	telecom *devices.Telecom
}

func (c *DisconnectTelecomCommand) Execute() {
	c.telecom.Disconnect()
}

func (c *DisconnectTelecomCommand) Undo() {
	c.telecom.Connect()
}

func (c *DisconnectTelecomCommand) GetName() string {
	return "DisconnectTelecom"
}

func NewDisconnectTelecomCommand(telecom *devices.Telecom) interfaces.Command {
	return &DisconnectTelecomCommand{telecom: telecom}
}

var _ interfaces.Command = &DisconnectTelecomCommand{}
