package entities

import (
	"github/elliot9/class12/infra"
	"github/elliot9/class12/internal/core"
	"io"
	"strconv"
	"strings"
)

const (
	// row indices
	HeaderLine = 1
	// column indices
	EmployeeId   = 0
	EmployeeName = 1
	EmployeeAge  = 2
	// optional column indices
	EmployeeSubordinates = 3
)

type RealDatabase struct {
	persistence *infra.Persistence
}

func NewRealDatabase(dbPath string) *RealDatabase {
	return &RealDatabase{
		persistence: infra.NewPersistence(dbPath),
	}
}

func (r *RealDatabase) GetEmployeeById(id int) (core.Employee, error) {
	line, err := r.persistence.ReadLine(id + HeaderLine)

	if err != nil {
		if err == io.EOF {
			return nil, &core.EmployeeNotFoundError{}
		}
		return nil, err
	}

	parts := strings.Split(line, " ")
	name := parts[EmployeeName]
	age, _ := strconv.Atoi(parts[EmployeeAge])
	subordinateIds := []int{}

	if len(parts) > EmployeeSubordinates {
		subordinates := strings.Split(parts[EmployeeSubordinates], ",")
		for _, subordinate := range subordinates {
			subordinateId, _ := strconv.Atoi(subordinate)
			subordinateIds = append(subordinateIds, subordinateId)
		}
	}

	return NewRealEmployee(id, name, age, subordinateIds, r), nil
}

var _ core.Database = &RealDatabase{}
