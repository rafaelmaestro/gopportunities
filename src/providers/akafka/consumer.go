package akafka

import (
	"fmt"
	"log/slog"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
)


type AKafkaConsumer struct {
	consumer any
}


func NewKafkaConsumer(config *config.Config) (*AKafkaConsumer, error) {
	go func(){
		kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9091",
			"group.id": "gopportunities",
			"auto.offset.reset": "earliest",
		})

		if err != nil {
			slog.Error("Failed to start a kafka consumer", "error", err)
		}

		fmt.Println(kafkaConsumer)


	}()

	return &AKafkaConsumer{}, nil
}
