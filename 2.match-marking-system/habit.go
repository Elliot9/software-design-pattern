package main

type Habit struct {
	name [10]byte
}

func (h *Habit) GetName() [10]byte {
	return h.name
}
