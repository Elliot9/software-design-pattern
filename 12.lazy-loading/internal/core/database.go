package core

type Database interface {
	GetEmployeeById(id int) (Employee, error)
}

type EmployeeNotFoundError struct{}

func (e *EmployeeNotFoundError) Error() string {
	return "employee not found"
}
