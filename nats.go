package dnslogging

import (
	"fmt"

	stan "github.com/nats-io/go-nats-streaming"
)

// NatsConfig is the required configuration for setting up the nats connection
type NatsConfig struct {
	natsurl   string
	clusterid string
	clientid  string
}

// ProcessNatsConfig takes in a config map and looks for the appropriate arguments to build the client with
func ProcessNatsConfig(config map[string]string) (nc NatsConfig, err error) {
	// Parse the URL
	nc.natsurl = config["natsUrl"]
	if nc.natsurl == "" {
		nc.natsurl = stan.DefaultNatsURL
	}

	// Parse the clusterid
	nc.clusterid = config["clusterId"]
	if nc.clusterid == "" {
		return nc, fmt.Errorf("clusterId was missing from configuration")
	}

	// Parse the clientid
	nc.clientid = config["clientId"]
	if nc.clientid == "" {
		return nc, fmt.Errorf("clientId was missing from configuration")
	}

	// Return the result
	return nc, nil
}
