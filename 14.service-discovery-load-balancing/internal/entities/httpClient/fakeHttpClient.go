package httpclient

import (
	"fmt"
	"github/elliot9/class14/internal/core"
	"time"

	"golang.org/x/exp/rand"
)

type FakeHttpClient struct {
	BaseHttpClient
}

func NewFakeHttpClient() *FakeHttpClient {
	f := &FakeHttpClient{}
	f.BaseHttpClient = BaseHttpClient{
		send: f.handleRequest,
	}
	return f
}

func (f *FakeHttpClient) handleRequest(request *core.HttpRequest) core.HttpStatusCode {
	targetURL := f.getResolvedURL(request)

	// 50% return a timeout
	if f.getRandomTrueOrFalse() {
		fmt.Printf("[Fail] %s\n", targetURL)
		return core.HttpTimeout
	}

	fmt.Printf("[Success] %s\n", targetURL)
	return core.HttpStatusCodeOK
}

func (f *FakeHttpClient) getResolvedURL(request *core.HttpRequest) string {
	if len(request.ResolvedIP) != 0 {
		newURL := *request.Url
		newURL.Host = request.ResolvedIP[0].Address
		return newURL.String()
	}
	return request.Url.String()
}

func (f *FakeHttpClient) getRandomTrueOrFalse() bool {
	seed := time.Now().UnixNano()
	source := rand.NewSource(uint64(seed))
	r := rand.New(source)
	return r.Intn(2) == 0
}
