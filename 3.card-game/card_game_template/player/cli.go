package player

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type CLI interface {
	ReadLine() (string, error)
	ReadNumber() (int, error)
	ReadYesNo() (bool, error)
}

type ConsoleIO struct {
	reader *bufio.Reader
}

func NewConsoleIO() *ConsoleIO {
	return &ConsoleIO{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (c *ConsoleIO) ReadLine() (string, error) {
	input, err := c.reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func (c *ConsoleIO) ReadNumber() (int, error) {
	input, err := c.ReadLine()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(input)
}

func (c *ConsoleIO) ReadYesNo() (bool, error) {
	input, err := c.ReadLine()
	if err != nil {
		return false, err
	}
	input = strings.ToLower(input)
	return input == "y" || input == "yes", nil
}
