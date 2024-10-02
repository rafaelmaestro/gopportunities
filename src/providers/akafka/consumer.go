package akafka

import (
	"context"
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
	Handle 	  func(ctx context.Context, msg *AKafkaMessage) error
}

func  NewKafkaConsumer(pctx context.Context, config *config.Config, kafkaConsumerConfig *AKafkaConsumerConfig) error {
	sLog := logger.Get()


	kafkaConsumer := &AKafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{"localhost:9091"}, // TODO: change to config
			GroupID:   kafkaConsumerConfig.ConsumerGroup,
			Topic:     kafkaConsumerConfig.Topic,
			Partition: 0,
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
			CommitInterval: 1 * time.Second,
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
		message, err := kafkaConsumer.reader.ReadMessage(pctx)
		if err != nil {
			sLog.Errorf("failed to read message on topic %s, partition %d: %v", kafkaConsumerConfig.Topic, message.Partition, err)
			continue
		}

		msg := &AKafkaMessage{
			AMessage: &message,
		}

		kafkaConsumerConfig.Handle(pctx, msg)
	}
}

// func NewKafkaConsumer(config *config.Config, consumerParams *AKafkaConsumer) error {
// 	go func(){
// 		kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
// 			"bootstrap.servers": "localhost:9091",
// 			"group.id": consumerParams.ConsumerGroup,
// 			"auto.offset.reset": "earliest",
// 		})

// 		if err != nil {
// 			slog.Error("Failed to start a kafka consumer", "error", err)
// 		}

// 		err = kafkaConsumer.SubscribeTopics(consumerParams.Topics, rebalanceCallback)

// 		if err != nil {
// 			slog.Error("Failed to subscribe to topics", "error", err)
// 		}

// 		slog.Debug("Kafka consumer started")

// 			// Set up a channel to handle shutdown
// 		sigchan := make(chan os.Signal, 1)
// 		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

// 		run := true
// 		for run {
// 			select {
// 				case sig := <-sigchan:
// 					slog.Debug("Caught signal", "signal", sig)
// 					run = false
// 				default:
// 					ev := kafkaConsumer.Poll(100)

// 					if ev == nil {
// 						continue
// 					}

// 					switch e := ev.(type) {
// 					case *kafka.Message:
// 						// fmt.Printf("Consumed Message: %s\n", string(e.Value))
// 						for _, topic := range consumerParams.Topics {
// 							if topic == *e.TopicPartition.Topic {
// 								fmt.Printf("Consumed Message: %s from  topic %s\n", string(e.Value), *e.TopicPartition.Topic)
// 							}
// 						}
// 					case kafka.Error:
// 						fmt.Printf("%% Error: %v: %v\n", e.Code(), e)
// 					default:
// 						// Handle other event types if necessary
// 					}
// 			Config
// 		}
// 	}()

// 	return nil
// }

// func rebalanceCallback(c *kafka.Consumer, event kafka.Event) error {
// 	switch ev := event.(type) {
// 	case kafka.AssignedPartitions:
// 		slog.Debug("[Kafka] Rebalance: new partitions assigned",
// 			"protocol", c.GetRebalanceProtocol(), "count", len(ev.Partitions), "partitions", ev.Partitions)

// 		// The application may update the start .Offset of each assigned
// 		// partition and then call Assign(). It is optional to call Assign
// 		// in case the application is not modifying any start .Offsets. In
// 		// that case we don't, the library takes care of it.
// 		// It is called here despite not modifying any .Offsets for illustrative
// 		// purposes.
// 		err := c.Assign(ev.Partitions)
// 		if err != nil {
// 			return err
// 		}

// 	case kafka.RevokedPartitions:
// 		slog.Debug("[Kafka] %% %s rebalance: %d partition(s) revoked",
// 			"protocol", c.GetRebalanceProtocol(), "count", len(ev.Partitions), "partitions", fmt.Sprintf("%v", ev.Partitions))

// 		// Usually, the rebalance callback for `RevokedPartitions` is called
// 		// just before the partitions are revoked. We can be certain that a
// 		// partition being revoked is not yet owned by any other consumer.
// 		// This way, logic like storing any pending offsets or committing
// 		// offsets can be handled.
// 		// However, there can be cases where the assignment is lost
// 		// involuntarily. In this case, the partition might already be owned
// 		// by another consumer, and operations including committing
// 		// offsets may not work.
// 		if c.AssignmentLost() {
// 			// Our consumer has been kicked out of the group and the
// 			// entire assignment is thus lost.
// 			slog.Error("[Kafka] Assignment lost involuntarily, commit may fail")
// 		}
// 		// Similar to Assign, client automatically calls Unassign() unless the
// 		// callback has already called that method. Here, we don't call it.

// 	default:
// 		slog.Debug("[Kafka] Unexpected event type: %v\n", event)
// 	}

// 	return nil
// }
