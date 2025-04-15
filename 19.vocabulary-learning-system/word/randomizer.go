package word

type Randomizer interface {
	Shuffle(n int, swap func(i, j int))
	IntN(n int) int
}
