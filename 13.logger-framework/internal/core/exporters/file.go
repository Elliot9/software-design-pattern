package exporters

import (
	"github/elliot9/class13/internal/core"
	"os"
)

type FileExporter struct {
	filename string
}

var _ core.Exporter = &FileExporter{}

func (e *FileExporter) Export(message string) {
	file, _ := os.OpenFile(e.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	file.WriteString(message + "\n")
}

func NewFileExporter(filename string) *FileExporter {
	return &FileExporter{filename: filename}
}
