package service

import (
	"bookcase_log/internal/storage"
	"time"

	"github.com/IBM/sarama"
)

type ServiceInterface interface {
	AddLog(msg *sarama.ConsumerMessage, ts time.Time) error
}

type Service struct {
	ServiceInterface
}

func New(storage storage.StorageInterface) *Service {
	return &Service{
		ServiceInterface: NewService(storage),
	}
}
