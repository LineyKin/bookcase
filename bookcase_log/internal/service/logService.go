package service

import (
	"bookcase_log/internal/storage"
	"bookcase_log/models"
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"
)

type logService struct {
	storage storage.StorageInterface
}

func NewService(s storage.StorageInterface) *logService {
	return &logService{
		storage: s,
	}
}

func (s *logService) AddLog(msg *sarama.ConsumerMessage) error {

	var pd models.Producerdata
	err := json.Unmarshal(msg.Value, &pd)
	if err != nil {
		log.Println("service AddLog error: ", err)
	}

	lr := models.LogRow{
		Producer_ts: pd.Timestamp,
		Consumer_ts: time.Now(),
		Topic:       msg.Topic,
		Message:     pd.Message,
	}

	return s.storage.AddLog(lr)
}
