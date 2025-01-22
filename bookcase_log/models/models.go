package models

import "time"

type LogRow struct {
	Producer_ts time.Time `json:"producer_ts"`
	Consumer_ts time.Time `json:"consumer_ts,omitempty"`
	Topic       string    `json:"topic,omitempty"`
	Message     string    `json:"message"`
}

type Producerdata struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
