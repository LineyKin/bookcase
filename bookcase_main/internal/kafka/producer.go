package kafka

import (
	"bookcase/lib/env"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type Producer struct {
	object sarama.SyncProducer
}

func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokers, config)
}

func getBrokers() []string {
	return []string{
		fmt.Sprintf("kafka:%s", env.GetKafkaPort()),
	}
}

func NewProducer() *Producer {
	p, err := ConnectProducer(getBrokers())
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
