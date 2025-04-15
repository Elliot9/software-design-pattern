package entities

import "github/elliot9/class12/internal/core"

type RealEmployee struct {
	Id           int
	Name         string
	Age          int
	Subordinates []core.Employee
}

func NewRealEmployee(id int, name string, age int, subordinateIds []int, database core.Database) *RealEmployee {
	subordinates := make([]core.Employee, 0)
	for _, id := range subordinateIds {
		subordinate, err := database.GetEmployeeById(id)
		if err != nil {
			if _, ok := err.(*core.EmployeeNotFoundError); ok {
				continue
			} else {
				panic(err)
			}
		}
		subordinates = append(subordinates, subordinate)
	}

	return &RealEmployee{
		Id:           id,
		Name:         name,
		Age:          age,
		Subordinates: subordinates,
	}
}

func (e *RealEmployee) GetId() int {
	return e.Id
}

func (e *RealEmployee) GetName() string {
	return e.Name
}

func (e *RealEmployee) GetAge() int {
	return e.Age
}

func (e *RealEmployee) GetSubordinates() []core.Employee {
	return e.Subordinates
}

var _ core.Employee = &RealEmployee{}
