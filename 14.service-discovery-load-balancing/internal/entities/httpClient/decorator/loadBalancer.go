package decorator

import (
	"github/elliot9/class14/internal/core"
)

type LoadBalancerDecorator struct {
	BaseDecorator
	ipIndexMap map[string]int
}

func NewLoadBalancerDecorator(next core.HttpClient) *LoadBalancerDecorator {
	l := &LoadBalancerDecorator{}
	l.BaseDecorator = BaseDecorator{
		next:       next,
		HttpClient: l,
	}
	l.initializeIPIndexMap()
	return l
}

func (l *LoadBalancerDecorator) initializeIPIndexMap() {
	l.ipIndexMap = make(map[string]int)
}

func (l *LoadBalancerDecorator) SendRequest(request *core.HttpRequest) core.HttpStatusCode {
	host, ips := request.Url.Host, request.ResolvedIP
	if len(ips) == 0 {
		return l.next.SendRequest(request)
	}

	ip := l.getNextIP(host, ips)
	request.ResolvedIP = []*core.IP{ip}

	return l.next.SendRequest(request)
}

func (l *LoadBalancerDecorator) getNextIP(host string, ips []*core.IP) *core.IP {
	if _, ok := l.ipIndexMap[host]; !ok {
		l.ipIndexMap[host] = 0
	}

	currentIndex := l.ipIndexMap[host]
	nextIndex := (currentIndex + 1) % len(ips)
	l.ipIndexMap[host] = nextIndex

	return ips[currentIndex%len(ips)]
}

var _ core.HttpClient = &LoadBalancerDecorator{}
