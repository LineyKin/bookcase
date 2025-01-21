package kafka

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

type Producer struct {
	object sarama.SyncProducer
}

const CONNECTION_TIME_LIMIT = 10

func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config = nil

	return sarama.NewSyncProducer(brokers, config)
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

func NewProducer() *Producer {
	brokers := getBrokers()
	p, err := ConnectProducer(brokers)

	startTimeStamp := time.Now()
	for err != nil && hasTimeToConnect(startTimeStamp) {
		log.Println("connection error: ", err)
		log.Println("connecting to kafka...")
		p, err = ConnectProducer(brokers)
	}

	if err != nil {
		log.Println("can't connect to kafka: ", err)
		return nil
	}

	kp := &Producer{
		object: p,
	}

	return kp
}

func (kp *Producer) Close() {
	kp.object.Close()
}
