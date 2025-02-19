package models

import (
	"time"
)

type UserLogInterface interface {
	NewLog() UserLog
}

type UserLog struct {
	Id        int       `json:"user_id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
