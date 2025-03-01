package models

import (
	"time"
)

type UserLogInterface interface {
	NewLog() UserLog
}

type UserLog struct {
	Id        int       `json:"user_id"`
	Login     string    `json:"login"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func NewUserLog() UserLog {
	var ul UserLog
	ul.Timestamp = time.Now()

	return ul
}
