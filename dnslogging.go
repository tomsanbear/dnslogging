package dnslogging

import (
	"context"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/miekg/dns"
	nats "github.com/nats-io/nats.go"
)

var log = clog.NewWithPlugin("dnslogging")

// DNSLogging is our plugin struct. No configurations being persisted at this time
type DNSLogging struct {
	nc      *NatsClient
	natsURL string

	Next plugin.Handler
}

// New creates a new instance of the DNSLogging type
func New() (*DNSLogging, error) {
	return &DNSLogging{
		natsURL: nats.DefaultURL,
	}, nil
}

func (dl *DNSLogging) Initialize() (err error) {
	dl.nc, err = NewNatsClient(dl.natsURL)
	if err != nil {
		return err
	}
	return nil
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
