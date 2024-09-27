package akafka

import (
	"log/slog"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
)

// TODO: refactor to use configs and all the other stuff

type IKafkaProducer interface {
	SendMessage(topic string, message string, key string) error
}

type AKafkaProducer struct  {
	producer *kafka.Producer
}

func NewKafkaProducer(config *config.Config) (*AKafkaProducer, error) {
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9091",
		"client.id": "gopportunities",
		"acks": "all",
		"batch.num.messages": 10,
		"queue.buffering.max.messages": 100000,
		"retries": 10,

	})


	if err != nil {
		slog.Error("Failed to start a kafka producer", "error", err)
		os.Exit(1)
		return nil, err
	}

	return &AKafkaProducer{
		producer: kafkaProducer,
	}, nil
}

func (p *AKafkaProducer) SendMessage(topic string, message string, key string) error {

	deliveryChan := make(chan kafka.Event, 100000)

	err := p.producer.Produce(&kafka.Message{
		Timestamp: time.Now(),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: []byte(message),
	},
		deliveryChan,
	)

	if err != nil {
		slog.Error("Failed to produce message", "error", err)
		return err
	}

	e := <-deliveryChan

	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		slog.Error("Delivery failed", "error", m.TopicPartition.Error)
		return m.TopicPartition.Error
	} else {
		// TODO: change to debug
		slog.Info("Delivered message to topic", "topic", *m.TopicPartition.Topic)
	}

	return nil
}
