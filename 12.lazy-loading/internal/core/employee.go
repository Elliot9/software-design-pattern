package core

type Employee interface {
	GetSubordinates() []Employee
	GetId() int
	GetName() string
	GetAge() int
}
