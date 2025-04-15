package decorator

import (
	"github/elliot9/class14/conf"
	"github/elliot9/class14/internal/core"
	"time"
)

const (
	FailDuration = 10 * time.Minute
)

type ServiceDiscoveryDecorator struct {
	BaseDecorator
	serverAddress map[string][]*core.IP
}

func NewServiceDiscoveryDecorator(next core.HttpClient) *ServiceDiscoveryDecorator {
	s := &ServiceDiscoveryDecorator{}
	s.BaseDecorator = BaseDecorator{
		next:       next,
		HttpClient: s,
	}
	s.initializeServerAddressFromConfig()
	return s
}

func (s *ServiceDiscoveryDecorator) initializeServerAddressFromConfig() {
	serviceDiscovery := conf.GetServiceDiscovery()
	s.serverAddress = make(map[string][]*core.IP)
	for host, ips := range serviceDiscovery {
		for _, ip := range ips {
			s.serverAddress[host] = append(s.serverAddress[host], &core.IP{Address: ip, IsActive: true})
		}
	}
}

func (d *ServiceDiscoveryDecorator) SendRequest(request *core.HttpRequest) core.HttpStatusCode {
	request.ResolvedIP = d.findActiveIPs(request.Url.Host)

	response := d.next.SendRequest(request)

	if response.IsFailure() {
		d.markIPAsFailed(request)
	}

	return response
}

func (d *ServiceDiscoveryDecorator) markIPAsFailed(request *core.HttpRequest) {
	if len(request.ResolvedIP) == 0 {
		return
	}

	host, targetIP := request.Url.Host, request.ResolvedIP[0]
	now := time.Now()

	if _, ok := d.serverAddress[host]; !ok {
		d.serverAddress[host] = []*core.IP{targetIP}
	}

	for _, ip := range d.serverAddress[host] {
		if ip == targetIP {
			ip.IsActive = false
			ip.LastFailTime = now
			return
		}
	}
}

func (d *ServiceDiscoveryDecorator) findActiveIPs(host string) []*core.IP {
	ips := []*core.IP{}
	for _, ip := range d.serverAddress[host] {
		if ip.IsActive {
			ips = append(ips, ip)
		} else {
			if ip.LastFailTime.IsZero() {
				ip.LastFailTime = time.Now()
			}
			if time.Since(ip.LastFailTime) > FailDuration {
				ip.IsActive = true
				ips = append(ips, ip)
			}
		}
	}
	return ips
}
