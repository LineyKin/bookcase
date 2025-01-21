package models

import "time"

type LogRow struct {
	Producer_ts time.Time
	Consumer_ts time.Time
	Topic       string
	Message     string
}

type Producerdata struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
