package guards

type NumberGuard func(number int) bool

func Equal(num int) NumberGuard {
	return func(n int) bool {
		return n == num
	}
}

func Less(number int) NumberGuard {
	return func(n int) bool {
		return n < number
	}
}

func Greater(number int) NumberGuard {
	return func(n int) bool {
		return n > number
	}
}

func GreaterAndEqual(number int) NumberGuard {
	return func(n int) bool {
		return n >= number
	}
}
