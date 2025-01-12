package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func (kp *Producer) PushLogToQueue(topic string, message []byte) error {

	// Create a new message
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Send message
	partition, offset, err := kp.object.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Author is stored in topic(%s)/partition(%d)/offset(%d)\n",
		topic,
		partition,
		offset)

	return nil
}
