package dnslogging

import (
	"context"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/miekg/dns"
	nats "github.com/nats-io/nats.go"
	"github.com/tomsanbear/recorder"
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

// Initialize handles the initial connection to the nats server.
func (dl *DNSLogging) Initialize() (err error) {
	log.Info("initializing the nats client")
	dl.nc, err = NewNatsClient(dl.natsURL)
	if err != nil {
		return err
	}
	log.Info("nats client was initialized")
	return nil
}

// Name is the name of our plugin
func (dl *DNSLogging) Name() string { return "dnslogging" }

// ServeDNS is doing the forwarding, this call can, and will fail since DNS resolution to the client should have succeeded by now
func (dl *DNSLogging) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (rc int, err error) {
	// This plugin is only valid when used in conjunction with the recorder plugin provided at github.com/tomsanbear/recorder
	recw := w.(*recorder.RecorderWriter)
	err = dl.nc.Publish(r, recw.Msg())
	if err != nil {
		return rc, err
	}
	return rc, err
}

// Close will drain and close of the nats connection
func (dl *DNSLogging) Close() {
	// Drain the connection
	clog.Info("draining the Nats client")
	dl.nc.conn.Drain()

	// Close the connection
	clog.Info("closing the Nats client")
	dl.nc.conn.Close()
}
