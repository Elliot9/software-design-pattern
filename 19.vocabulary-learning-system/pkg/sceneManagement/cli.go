package sceneManagement

type CLI interface {
	Println(s string)
	ReadLine() string
}
