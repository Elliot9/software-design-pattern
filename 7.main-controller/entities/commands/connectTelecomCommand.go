package commands

import (
	"github/elliot9/class7/entities/devices"
	"github/elliot9/class7/interfaces"
)

type ConnectTelecomCommand struct {
	telecom *devices.Telecom
}

func (c *ConnectTelecomCommand) Execute() {
	c.telecom.Connect()
}

func (c *ConnectTelecomCommand) Undo() {
	c.telecom.Disconnect()
}

func (c *ConnectTelecomCommand) GetName() string {
	return "ConnectTelecom"
}

func NewConnectTelecomCommand(telecom *devices.Telecom) interfaces.Command {
	return &ConnectTelecomCommand{telecom: telecom}
}

var _ interfaces.Command = &ConnectTelecomCommand{}
