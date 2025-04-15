package httpclient

import (
	"github/elliot9/class14/internal/core"
)

type BaseHttpClient struct {
	send func(*core.HttpRequest) core.HttpStatusCode
}

func (h *BaseHttpClient) SendRequest(request *core.HttpRequest) core.HttpStatusCode {
	// if the request has a resolved IP, use the first one as default
	if len(request.ResolvedIP) != 0 {
		request.ResolvedIP = []*core.IP{request.ResolvedIP[0]}
	}

	return h.send(request)
}

var _ core.HttpClient = &BaseHttpClient{}
