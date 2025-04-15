package devices

import "fmt"

type Telecom struct{}

func (t *Telecom) Connect() {
	fmt.Println("The telecom has been turned on.")
}

func (t *Telecom) Disconnect() {
	fmt.Println("The telecom has been turned off.")
}

func NewTelecom() *Telecom {
	return &Telecom{}
}
