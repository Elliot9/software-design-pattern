package main

import (
	"fmt"
	"github/elliot9/class12/infra"
	"github/elliot9/class12/internal/entities"
	"github/elliot9/class12/internal/proxy"
)

func main() {
	realDatabase := entities.NewRealDatabase("employee.txt")
	virtualRealDatabaseProxy := proxy.NewVirtualRealDatabaseProxy(realDatabase)
	passwordProtectionDatabaseProxy := proxy.NewPasswordProtectionDatabaseProxy(infra.NewEnvReader(), virtualRealDatabaseProxy)

	employee, err := passwordProtectionDatabaseProxy.GetEmployeeById(2)
	if err != nil {
		panic(err)
	}

	subordinates := employee.GetSubordinates()
	for _, subordinate := range subordinates {
		fmt.Println(subordinate.GetName())
	}
}
