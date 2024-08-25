package aggregator

import (
	"sync"
	"time"
	"trading-app-service/internal/models"
)

type Aggregator struct {
	mu      sync.Mutex
	candles map[string]*models.Candle
}

func NewAggregator() *Aggregator {
	return &Aggregator{
		candles: make(map[string]*models.Candle),
	}
}

func (a *Aggregator) ProcessTick(symbol string, price float64, volume float64) *models.Candle {
	a.mu.Lock()
	defer a.mu.Unlock()

	now := time.Now().Truncate(time.Minute)

	candle, exists := a.candles[symbol]
	if !exists || candle.Timestamp.Before(now) {
		if exists {
			completeCandle := *candle
			a.candles[symbol] = &models.Candle{Timestamp: now, Open: price, High: price, Low: price, Close: price, Volume: volume}
			return &completeCandle
		}
		a.candles[symbol] = &models.Candle{Timestamp: now, Open: price, High: price, Low: price, Close: price, Volume: volume}
		return nil
	}

	candle.High = max(candle.High, price)
	candle.Low = min(candle.Low, price)
	candle.Close = price
	candle.Volume += volume

	return nil
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
