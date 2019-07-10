package dnslogging

import "github.com/coredns/coredns/request"

// Client is the interface that other stremaing clients must implement to be used.
type Client interface {
	Publish(request.Request) error
}
