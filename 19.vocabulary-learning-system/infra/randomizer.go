package infra

import "math/rand/v2"

type RealRandomizer struct{}

func NewRealRandomizer() *RealRandomizer {
	return &RealRandomizer{}
}

func (r *RealRandomizer) Shuffle(n int, swap func(i, j int)) {
	rand.Shuffle(n, swap)
}

func (r *RealRandomizer) IntN(n int) int {
	return rand.IntN(n)
}

type MockRandomizer struct{}

func NewMockRandomizer() *MockRandomizer {
	return &MockRandomizer{}
}

func (r *MockRandomizer) Shuffle(n int, swap func(i, j int)) {
}

func (r *MockRandomizer) IntN(n int) int {
	return 0
}
