package consumer

import (
	"bookcase_log/lib/env"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

const TOPIC = "bookcase_log"

type KafkaConsumer struct {
	consumer sarama.Consumer
}

func connectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config = nil

	return sarama.NewConsumer(brokers, config)
}

func New() (*KafkaConsumer, error) {
	broker := fmt.Sprintf("kafka:%s", env.GetKafkaPort())
	worker, err := connectConsumer([]string{broker})
	if err != nil {
		log.Println("can't connect to kafka consumer", err)
		return nil, err
	}

	kc := &KafkaConsumer{
		consumer: worker,
	}

	return kc, nil
}

func (kc *KafkaConsumer) Partition() (sarama.PartitionConsumer, error) {
	return kc.consumer.ConsumePartition(TOPIC, 0, sarama.OffsetOldest)
}

func (kc *KafkaConsumer) Close() error {
	return kc.consumer.Close()
}
