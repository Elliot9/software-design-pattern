package proxy

import (
	"github/elliot9/class12/internal/core"
	"github/elliot9/class12/internal/entities"
)

type VirtualRealEmployeeProxy struct {
	employee *entities.RealEmployee
	database core.Database
	id       int
}

func NewVirtualRealEmployeeProxy(id int, database core.Database) *VirtualRealEmployeeProxy {
	return &VirtualRealEmployeeProxy{
		employee: nil,
		database: database,
		id:       id,
	}
}

func (p *VirtualRealEmployeeProxy) lazyInitialization() error {
	if p.employee == nil {
		employee, err := p.database.GetEmployeeById(p.id)
		if err != nil {
			return err
		}
		p.employee = employee.(*entities.RealEmployee)
	}
	return nil
}

func (v *VirtualRealEmployeeProxy) GetId() int {
	err := v.lazyInitialization()
	if err != nil {
		panic(err)
	}
	return v.employee.GetId()
}

func (v *VirtualRealEmployeeProxy) GetName() string {
	err := v.lazyInitialization()
	if err != nil {
		panic(err)
	}
	return v.employee.GetName()
}

func (v *VirtualRealEmployeeProxy) GetAge() int {
	err := v.lazyInitialization()
	if err != nil {
		panic(err)
	}
	return v.employee.GetAge()
}

func (v *VirtualRealEmployeeProxy) GetSubordinates() []core.Employee {
	err := v.lazyInitialization()
	if err != nil {
		panic(err)
	}
	return v.employee.GetSubordinates()
}

var _ core.Employee = &VirtualRealEmployeeProxy{}
