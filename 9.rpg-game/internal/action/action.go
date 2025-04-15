package action

import "github/elliot9/class9/internal/interfaces"

type BaseAction struct {
	interfaces.Action
	Name        string
	MPCost      int
	TargetCount int
	TargetType  interfaces.TargetType
}

func (a *BaseAction) GetName() string {
	return a.Name
}

func (a *BaseAction) GetMPCost() int {
	return a.MPCost
}

func (a *BaseAction) GetTargetCount() int {
	return a.TargetCount
}

func (a *BaseAction) GetTargetType() interfaces.TargetType {
	return a.TargetType
}

func (a *BaseAction) Excute(role interfaces.Role, targets []interfaces.Role) {
	a.Action.Excute(role, targets)
}

var _ interfaces.Action = &BaseAction{}
