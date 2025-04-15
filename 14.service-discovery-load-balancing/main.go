package main

import (
	"fmt"
	"github/elliot9/class14/internal/core"
	httpclient "github/elliot9/class14/internal/entities/httpClient"
	"github/elliot9/class14/internal/entities/httpClient/decorator"
)

func main() {
	httpClient := decorator.NewBlackListDecorator(decorator.NewServiceDiscoveryDecorator(decorator.NewLoadBalancerDecorator(httpclient.NewFakeHttpClient())))

	for i := 0; i < 10; i++ {
		request := &core.HttpRequest{
			Url: &core.URL{
				Scheme: "http",
				Host:   "waterballsa.tw",
				Path:   "/mail",
			},
		}

		response := httpClient.SendRequest(request)
		fmt.Printf("response: %s\nstatus code: %d\n", response.String(), response)
	}
}
