package akafka

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/fnunezzz/go-logger"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	"github.com/segmentio/kafka-go"
)

type AKafkaMessage struct {
    AMessage *kafka.Message
}

type AKafkaConsumer struct {
	reader *kafka.Reader
}

type AKafkaConsumerConfig struct {
	ConsumerGroup string
	Topic       string
	ConcurrentReaders int // this config is useful here because we can set different number of concurrent readers for different topics
	MessageDto reflect.Type
	Handle 	  func(ctx context.Context, dto interface{} ) error
}

func NewKafkaConsumer(cfg *config.Config, kafkaConsumerConfig *AKafkaConsumerConfig) error {
	sLog := logger.Get()

	minBytes, err := strconv.Atoi(cfg.Kafka.ConsumerMinBytes); if err != nil {
		minBytes = 10e3
	}

	maxBytes, err := strconv.Atoi(cfg.Kafka.ConsumerMaxBytes); if err != nil {
		maxBytes = 10e6
	}

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	for i := 0; i < kafkaConsumerConfig.ConcurrentReaders; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			kafkaConsumer := &AKafkaConsumer{
				reader: kafka.NewReader(kafka.ReaderConfig{
					Brokers:   []string{cfg.Kafka.Brokers},
					GroupID:   kafkaConsumerConfig.ConsumerGroup,
					Topic:     kafkaConsumerConfig.Topic,
					MinBytes: minBytes, // 10KB
					MaxBytes: maxBytes, // 10MB
					HeartbeatInterval: time.Duration(cfg.Kafka.HeartbeatInterval) * time.Second,
					SessionTimeout: time.Duration(cfg.Kafka.HeartbeatInterval * cfg.Kafka.SessionTimeoutMultiplier) * time.Second,
					CommitInterval: 1 * time.Second, // TODO: change to config
					StartOffset: kafka.FirstOffset,
					ErrorLogger: kafka.LoggerFunc(sLog.Errorf),
				}),
			}

			defer func() {
				if err := kafkaConsumer.reader.Close(); err != nil {
					sLog.Errorf("failed to close kafka reader: %v", err)
				}
			}()

			for {
				message, err := kafkaConsumer.reader.ReadMessage(ctx)
				if err != nil {
					sLog.Errorf("failed to read message on topic %s, partition %d: %v", kafkaConsumerConfig.Topic, message.Partition, err)
					continue
				}

                dtoInstance := reflect.New(kafkaConsumerConfig.MessageDto).Interface()
                if err := json.Unmarshal(message.Value, dtoInstance); err != nil {
                    sLog.Errorf("failed to unmarshal message: %v", err)
                    continue
                }

                if err := kafkaConsumerConfig.Handle(ctx, dtoInstance); err != nil {
                    sLog.Errorf("error handling message: %v", err)
                }
			}
		}()
	}
	wg.Wait()
	cancel()
	return nil
}
