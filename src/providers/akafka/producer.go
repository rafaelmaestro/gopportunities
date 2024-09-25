package akafka

import (
	"fmt"
	"log/slog"

	"github.com/IBM/sarama"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
)

type IKafkaProducer interface {
	SendMessage(topic string, message string, key string) error
	Close() error
}

type AKafkaProducer struct  {
	producer sarama.SyncProducer
}

func NewKafkaProducer(config *config.Config) (*AKafkaProducer, error) {
	kafkaConnectionRetries := config.Kafka.ConnectionRetries


	fmt.Println(config.Kafka.Brokers)

	brokers := []string{"localhost:9091"}

	fmt.Println(brokers)

	kafkaConfig := sarama.NewConfig()

	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Retry.Max = config.Kafka.ProducerRetries


	var kafkaProducer sarama.SyncProducer
	var err error

	for kafkaConnectionRetries > 0 {
		kafkaProducer, err = sarama.NewSyncProducer(brokers, kafkaConfig)

		if err != nil {
			slog.Error("Failed to start a kafka producer", "error", err)
			slog.Error("Retrying to connect to kafka... ", "retries", kafkaConnectionRetries)
			kafkaConnectionRetries--
		} else {
			break
		}

		if kafkaConnectionRetries == 0 {
			return nil, err
		}
	}

	return &AKafkaProducer{
		producer: kafkaProducer,
	}, nil
}

func (p *AKafkaProducer) SendMessage(topic string, message string, key string) error {
	kafkaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Key: sarama.StringEncoder(key),
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := p.producer.SendMessage(kafkaMessage)

	if err != nil {
		slog.Error("Failed to send message to kafka", "error", err)
		return err
	}

	slog.Debug("Message is stored in topic/partition/offset", "topic", topic, "partition", partition, "offset", offset)

	return nil
}

func (p *AKafkaProducer) Close() error {
	return p.producer.Close()
}
