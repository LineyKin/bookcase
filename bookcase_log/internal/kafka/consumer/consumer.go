package consumer

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

const TOPIC = "bookcase_log"
const CONNECTION_TIME_LIMIT = 10

type KafkaConsumer struct {
	consumer sarama.Consumer
}

func connectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config = nil

	return sarama.NewConsumer(brokers, config)
}

func getBrokers() []string {
	return []string{
		"kafka:9092",
	}
}

func hasTimeToConnect(startTimeStamp time.Time) bool {
	now := time.Now()
	dur := now.Sub(startTimeStamp)
	log.Println("connecting time: ", dur)

	return int(dur.Seconds()) < CONNECTION_TIME_LIMIT
}

func New() (*KafkaConsumer, error) {
	brokers := getBrokers()
	worker, err := connectConsumer(brokers)
	startTimeStamp := time.Now()
	for err != nil && hasTimeToConnect(startTimeStamp) {
		log.Println("connection error: ", err)
		log.Println("connecting to kafka...")
		worker, err = connectConsumer(brokers)
	}

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
