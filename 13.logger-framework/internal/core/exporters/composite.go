package exporters

import (
	"github/elliot9/class13/internal/core"
)

type CompositeExporter struct {
	exporters []core.Exporter
}

var _ core.Exporter = &CompositeExporter{}

func (e *CompositeExporter) Export(message string) {
	for _, exporter := range e.exporters {
		exporter.Export(message)
	}
}

func NewCompositeExporter(exporters []core.Exporter) *CompositeExporter {
	return &CompositeExporter{exporters: exporters}
}
