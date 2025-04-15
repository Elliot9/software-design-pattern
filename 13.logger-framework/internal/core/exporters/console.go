package exporters

import (
	"fmt"
	"github/elliot9/class13/internal/core"
)

type ConsoleExporter struct{}

var _ core.Exporter = &ConsoleExporter{}

func (e *ConsoleExporter) Export(message string) {
	fmt.Println(message)
}

func NewConsoleExporter() *ConsoleExporter {
	return &ConsoleExporter{}
}
