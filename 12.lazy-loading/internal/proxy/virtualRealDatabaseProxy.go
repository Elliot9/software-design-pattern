package proxy

import (
	"github/elliot9/class12/internal/core"
	"github/elliot9/class12/internal/entities"
)

type VirtualRealDatabaseProxy struct {
	database *entities.RealDatabase
}

func NewVirtualRealDatabaseProxy(database *entities.RealDatabase) *VirtualRealDatabaseProxy {
	return &VirtualRealDatabaseProxy{
		database: database,
	}
}

func (v *VirtualRealDatabaseProxy) GetEmployeeById(id int) (core.Employee, error) {
	return NewVirtualRealEmployeeProxy(id, v.database), nil
}

var _ core.Database = &VirtualRealDatabaseProxy{}
