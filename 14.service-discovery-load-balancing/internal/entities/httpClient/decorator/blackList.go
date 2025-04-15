package decorator

import (
	"github/elliot9/class14/conf"
	"github/elliot9/class14/internal/core"
)

type BlackListDecorator struct {
	BaseDecorator
	blackList map[string]bool
}

func NewBlackListDecorator(next core.HttpClient) *BlackListDecorator {
	b := &BlackListDecorator{}
	b.BaseDecorator = BaseDecorator{
		next:       next,
		HttpClient: b,
	}
	b.initializeBlackList()
	return b
}

func (b *BlackListDecorator) initializeBlackList() {
	b.blackList = conf.GetBlackList()
}

func (b *BlackListDecorator) SendRequest(request *core.HttpRequest) core.HttpStatusCode {
	if b.blackList[request.Url.Host] {
		return core.HttpStatusCodeForbidden
	}
	return b.next.SendRequest(request)
}

var _ core.HttpClient = &BlackListDecorator{}
