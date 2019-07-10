package dnslogging

import (
	retry "github.com/avast/retry-go"
	"github.com/coredns/coredns/request"
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

func (nc *NatsClient) Publish(state request.Request) error {
	// Get our objects we need
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return err
	}
	data, err := dnsproto.NewRootData(seg)
	if err != nil {
		return err
	}
	req, err := data.NewRequest()
	if err != nil {
		return err
	}
	questions, err := req.NewQuestion(int32(len(state.Req.Question)))
	for i, question := range state.Req.Question {
		questions.Set(i, question.String())
	}
	resp, err := data.NewResponse()
	if err != nil {
		return err
	}
	answers, err := resp.NewAnswers(int32(len(state.Resp.Answer)))
	if err != nil {
		return err
	}
	for i, answer := range state.Resp.Answer {
		answers.Set(i, answer.String())
	}

	// Publish the bytes to the wire
	mrshled, err := msg.Marshal()
	if err != nil {
		return err
	}
	err = nc.conn.Publish(nc.subject, mrshled)

	return nil
}
