package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CLI interface {
	Println(s string)
	ReadLine() string
	ReadNumber() []int
}

type ConsoleIO struct {
	reader *bufio.Reader
}

func NewConsoleIO() *ConsoleIO {
	return &ConsoleIO{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (c *ConsoleIO) ReadLine() string {
	input, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (c *ConsoleIO) ReadNumber() []int {
	input := c.ReadLine()
	numbers := []int{}
	for _, number := range strings.Split(input, ", ") {
		num, err := strconv.Atoi(number)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func (c *ConsoleIO) Println(s string) {
	fmt.Println(s)
}

var _ CLI = &ConsoleIO{}
