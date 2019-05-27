package dnslogging

import (
	"fmt"
	"time"

	stan "github.com/nats-io/go-nats-streaming"
)

// UpstreamClientType is the constant for stuff
type UpstreamClientType string

// The supported types we have
const (
	nats UpstreamClientType = "nats"
)

// UpstreamClient represents an interface to any of the supported clients, since multiple upstreams can be defined
type UpstreamClient struct {
	upstreamtype UpstreamClientType
	upstream     interface{}
}

// NewUpstreamClient takes matches the type provided by the config and initializes a connection with the appropriate configs
func NewUpstreamClient(upstreamtype UpstreamClientType, config map[string]string) (uc *UpstreamClient, err error) {
	uc.upstreamtype = upstreamtype
	switch upstreamtype {
	case nats:
		natsConfig, err := ProcessNatsConfig(config)
		if err != nil {
			return uc, err
		}
		uc.upstream, err = stan.Connect(
			natsConfig.clusterid, natsConfig.clientid, stan.NatsURL(natsConfig.natsurl), stan.PubAckWait(time.Second*15),
		)
		if err != nil {
			return uc, err
		}
	default:
		return nil, fmt.Errorf("%v is not a supported type of upstream client", upstreamtype)
	}
	return uc, err
}

// Publish writes the raw string as an event for the source specified
func (uc *UpstreamClient) Publish(event string) (err error) {
	switch uc.upstreamtype {
	case nats:
		natsClient := uc.upstream.(stan.Conn)
		if natsClient == nil {
			return fmt.Errorf("error creating Nats client")
		}
		err = natsClient.Publish("dnslogging", []byte(event))
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("publishing is unsupported for client: %v", uc.upstreamtype)
	}
	return nil
}
