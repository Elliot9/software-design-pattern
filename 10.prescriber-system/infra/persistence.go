package infra

import (
	"os"
)

const (
	FileMode = 0644 // "0644 (rw-r--r--)"
)

type Persistence interface {
	Load() string
	Save(output string)
	Replace(output string)
}

type FilePersistence struct {
	filePath string
}

func (p *FilePersistence) Load() string {
	jsonStr, err := os.ReadFile(p.filePath)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}

func (p *FilePersistence) Save(output string) {
	file, _ := os.OpenFile(p.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, FileMode)
	defer file.Close()

	file.WriteString(output + "\n")
}

func (p *FilePersistence) Replace(output string) {
	file, _ := os.OpenFile(p.filePath, os.O_CREATE|os.O_WRONLY, FileMode)
	defer file.Close()

	file.WriteString(output)
}

func NewFilePersistence(filePath string) *FilePersistence {
	return &FilePersistence{filePath: filePath}
}
