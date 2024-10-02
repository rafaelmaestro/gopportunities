package akafka

import (
	"context"
	"time"

	"github.com/fnunezzz/go-logger"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	"github.com/segmentio/kafka-go"
)

// TODO: refactor to use configs and all the other stuff

type IKafkaProducer interface {
	SendMessage(pctx context.Context, topic string, message string, key string, headers map[string]string) error
	Close() error
}

type AKafkaProducer struct  {
	producer *kafka.Writer
}

func NewKafkaProducer(config *config.Config) *AKafkaProducer {
	return &AKafkaProducer{
		producer: &kafka.Writer{
			Addr: kafka.TCP("localhost:9091"), // TODO: change to config
			Balancer: &kafka.LeastBytes{},
			BatchSize: 1000, // TODO: change to config
			BatchTimeout: 50 * time.Millisecond, // TODO: change to config
			MaxAttempts: config.Kafka.ProducerRetries ,
			RequiredAcks: kafka.RequireAll, // TODO: change to config
			Async: false, // TODO: change to config (DEFAULT: false)
			ErrorLogger: kafka.LoggerFunc(logger.Get().Errorf),
		},
	}
}

func (p *AKafkaProducer) SendMessage(pctx context.Context, topic string, message string, key string, headers map[string]string) error {
	headersConverted := convertHeaders(headers)

	err := p.producer.WriteMessages(
		pctx,
		kafka.Message{
			Topic: topic,
			Time: time.Now(),
			Key: []byte(key),
			Value: []byte(message),
			Headers: headersConverted,
		},
	)

	if err != nil {
		return err
	}
	return nil
}

func convertHeaders(headers map[string]string) []kafka.Header {
	var kafkaHeaders []kafka.Header
	for k, v := range headers {
		kafkaHeaders = append(kafkaHeaders, kafka.Header{Key: k, Value: []byte(v)})
	}
	return kafkaHeaders
}

func (p *AKafkaProducer) Close() error {
	return p.producer.Close()
}
