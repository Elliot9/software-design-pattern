package decorator

import (
	"github/elliot9/class14/internal/core"
)

type BaseDecorator struct {
	next core.HttpClient
	core.HttpClient
}

func (d *BaseDecorator) SendRequest(request *core.HttpRequest) core.HttpStatusCode {
	d.HttpClient.SendRequest(request)

	return d.next.SendRequest(request)
}

var _ core.HttpClient = &BaseDecorator{}
