package core

type Exporter interface {
	Export(message string)
}
