package dnslogging

import "github.com/miekg/dns"

// Client is the interface that other stremaing clients must implement to be used.
type Client interface {
	// Publish takes a request and publishes it to the streaming service, this should ideally be batching for performance, but that is left to programmer.
	Publish(req *dns.Msg, resp *dns.Msg) error
}
