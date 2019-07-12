package dnslogging

import (
	retry "github.com/avast/retry-go"
	"github.com/miekg/dns"
	nats "github.com/nats-io/nats.go"
	"github.com/tomsanbear/dnslogging/dnsproto"
	capnp "zombiezen.com/go/capnproto2"
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

func (nc *NatsClient) Publish(req *dns.Msg, resp *dns.Msg) error {
	// Get our objects we need
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return err
	}
	pdata, err := dnsproto.NewRootData(seg)
	if err != nil {
		return err
	}
	preq, err := pdata.NewRequest()
	if err != nil {
		return err
	}
	pquestions, err := preq.NewQuestion(int32(len(req.Question)))
	for i, question := range req.Question {
		pquestions.Set(i, question.String())
	}
	presp, err := pdata.NewResponse()
	if err != nil {
		return err
	}
	panswers, err := presp.NewAnswers(int32(len(resp.Answer)))
	if err != nil {
		return err
	}
	for i, answer := range resp.Answer {
		panswers.Set(i, answer.String())
	}

	// Publish the bytes to the wire
	mrshled, err := msg.Marshal() 
	if err != nil {
		return err
	}
	err = nc.conn.Publish(nc.subject, mrshled)

	return nil
}
 