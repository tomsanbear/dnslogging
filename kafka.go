package dnslogging

import (
	"context"

	retry "github.com/avast/retry-go"
	kafka "github.com/segmentio/kafka-go"
)

// KafkaClient contains a kafka writer which can be used to async send messages to kafka
type KafkaClient struct {
	writer *kafka.Writer
}

// NewKafkaClient creates a kafka client for the plugin to publish to
func NewKafkaClient(brokers []string, topic string) (kc *KafkaClient, err error) {
	err = retry.Do(
		func() error {
			kc.writer = kafka.NewWriter(kafka.WriterConfig{
				Brokers:  brokers,
				Topic:    topic,
				Balancer: &kafka.LeastBytes{},
			})
			if kc.writer == nil {
				return &Error{"failed to initialize kafka writer"}
			}
			return err
		},
		retry.DelayType(retry.BackOffDelay),
		retry.OnRetry(func(n uint, err error) {
			log.Errorf("kafka: (attempt %v) %v", n+1, err)
		}),
	)
	if err != nil {
		return nil, err
	}
	return kc, nil
}

// Publish sends a log object to kafka
func (kc *KafkaClient) Publish(log *Log) (err error) {
	logBytes, err := log.pack()
	if err != nil {
		return err
	}
	err = kc.writer.WriteMessages(context.Background(), kafka.Message{
		Value: logBytes,
	})
	if err != nil {
		return err
	}
	return nil
}
