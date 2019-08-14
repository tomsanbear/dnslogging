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
	// Nats configuration
	nc      *NatsClient
	natsURL string

	// Kafka configuration
	kc           *KafkaClient
	kafkaBrokers []string
	kafkaTopic   string

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
	// Nats initialization
	if dl.natsURL != "" {
		log.Info("initializing the nats client")
		dl.nc, err = NewNatsClient(dl.natsURL)
		if err != nil {
			return err
		}
	}

	// Kafka Initialization
	if dl.kafkaBrokers != nil && dl.kafkaTopic != "" {
		log.Info("initializing the kafka client")
		dl.kc, err = NewKafkaClient(dl.kafkaBrokers, dl.kafkaTopic)
		if err != nil {
			return err
		}
	}

	return nil
}

// Name is the name of our plugin
func (dl *DNSLogging) Name() string { return "dnslogging" }

// ServeDNS is doing the forwarding, this call can, and will fail since DNS resolution to the client should have succeeded by now
func (dl *DNSLogging) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (rc int, err error) {
	// This plugin is only valid when used in conjunction with the recorder plugin provided at github.com/tomsanbear/recorder
	// TODO: make this smart to recognize when we have access to the response object
	recw := w.(*recorder.RecorderWriter)
	if recw == nil {
		return rc, &Error{"invalid writer used, unable to extract response"}
	}

	// Formulate the log
	log := &Log{
		Req:  r,
		Resp: recw.Msg(),
	}

	// Publish to whatever available streaming services we have
	if dl.nc != nil {
		err = dl.nc.Publish(log)
		if err != nil {
			return rc, err
		}
	}
	if dl.kc != nil {
		err = dl.kc.Publish(log)
		if err != nil {
			return rc, err
		}
	}
	return rc, err
}

// Close will drain and close of the nats connection
func (dl *DNSLogging) Close() {
	// Drain the connection
	clog.Info("draining the Nats client")
	err := dl.nc.conn.Drain()
	if err != nil {
		log.Warning("nats: ", err)
	}

	// Close the connection
	clog.Info("closing the Nats client")
	dl.nc.conn.Close()
}
