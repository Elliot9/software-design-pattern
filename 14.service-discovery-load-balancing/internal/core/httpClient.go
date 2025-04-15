package core

import (
	"fmt"
	"net/url"
	"time"
)

type HttpClient interface {
	SendRequest(request *HttpRequest) HttpStatusCode
}

type HttpStatusCode int

const (
	HttpStatusCodeOK        HttpStatusCode = 200
	HttpStatusCodeForbidden HttpStatusCode = 403
	HttpTimeout             HttpStatusCode = 504
)

func (s HttpStatusCode) IsFailure() bool {
	return s != HttpStatusCodeOK
}

func (s HttpStatusCode) String() string {
	switch s {
	case HttpStatusCodeOK:
		return "OK"
	case HttpStatusCodeForbidden:
		return "Forbidden"
	case HttpTimeout:
		return "Timeout"
	default:
		return "Unknown"
	}
}

type HttpRequest struct {
	Url        *URL
	ResolvedIP []*IP
}

type URL struct {
	Scheme string
	Host   string
	Path   string
}

func (u *URL) String() string {
	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.Path)
}

type IP struct {
	Address      string
	IsActive     bool
	LastFailTime time.Time
}

func NewURL(targetURL string) *URL {
	parsedURL, err := url.Parse(targetURL)

	if err != nil {
		return nil
	}

	return &URL{
		Scheme: parsedURL.Scheme,
		Host:   parsedURL.Host,
		Path:   parsedURL.Path,
	}
}

func NewIP(address string) *IP {
	return &IP{
		Address: address,
	}
}
