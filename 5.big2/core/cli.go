package core

type CLI interface {
	ReadLine() string
	ReadNumber() []int
	Println(string)
}
