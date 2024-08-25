package models

import (
	"time"
)

type Candle struct {
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
}

type ChartCandle struct {
	Symbol         string
	Timestamp      time.Time `json:"timestamp"`
	Open           float64   `json:"open"`
	High           float64   `json:"high"`
	Low            float64   `json:"low"`
	Close          float64   `json:"close"`
	Volume         float64   `json:"volume"`
	MarkPoint      string    `json:"markPoint"`
	MarkPointPrice float64   `json:"markPointPrice"`
}
