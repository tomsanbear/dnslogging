package dnslogging

import (
	retry "github.com/avast/retry-go"
	nats "github.com/nats-io/nats.go"
)

type NatsClient struct {
	conn    *nats.Conn
	subject string
}

// NewNatsClient returns a connected Nats Client
func NewNatsClient(natsURL string) (_ *NatsClient, err error) {
	nc := NatsClient{}
	// Initial Connection loop
	err = retry.Do(
		func() error {
			nc.conn, err = nats.Connect(natsURL)
			return err
		},
		retry.DelayType(retry.BackOffDelay),
		retry.OnRetry(func(n uint, err error) {
			log.Errorf("nats: (attempt %v) %v", n+1, err)
		}),
	)
	if err != nil {
		return nil, err
	}
	log.Info("nats connection initialized...")
	return &nc, nil
}

func (nc *NatsClient) Publish(log *Log) error {
	lb, err := log.pack()
	if err != nil {
		return err
	}
	err = nc.conn.Publish(nc.subject, lb)
	if err != nil {
		return err
	}
	return nil
}
