package infra

import (
	"bufio"
	"io"
	"os"
)

type Persistence struct {
	filePath string
}

func (p *Persistence) ReadLine(row int) (line string, err error) {
	currentLine := 0

	reader, err := os.Open(p.filePath)
	if err != nil {
		return "", err
	}

	sc := bufio.NewScanner(reader)
	for sc.Scan() {
		currentLine++
		if currentLine == row {
			return sc.Text(), sc.Err()
		}
	}
	return line, io.EOF
}

func NewPersistence(filePath string) *Persistence {
	return &Persistence{
		filePath: filePath,
	}
}
