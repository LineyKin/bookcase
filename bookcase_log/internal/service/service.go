package service

import (
	"bookcase_log/internal/storage"
	"bookcase_log/models"
	"time"

	"github.com/IBM/sarama"
)

type ServiceInterface interface {
	AddLog(msg *sarama.ConsumerMessage, ts time.Time) error
	GetLogCount() (int, error)
	GetLogList(limit, offset int) ([]models.LogRow, error)
}

type Service struct {
	ServiceInterface
}

func New(storage storage.StorageInterface) *Service {
	return &Service{
		ServiceInterface: NewService(storage),
	}
}
