package dnslogging

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
)

// DNSLogging is our plugin struct. No configurations being persisted at this time
type DNSLogging struct {
	Next plugin.Handler
}

// New creates a new instance of the DNSLogging type
func New() (*DNSLogging, error) {
	return &DNSLogging{}, nil
}

// Name is the name of our plugin
func (dl *DNSLogging) Name() string { return "dnslogging" }

// ServeDNS is doing the forwarding, this call can, and will fail since DNS resolution to the client should have succeeded by now
func (dl *DNSLogging) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (rc int, err error) {
	return rc, err
}

func (dl *DNSLogging) Close() {
	// Close the connection to the server here
}
