package db

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"log"
	"time"
	"trading-app-service/internal/models"
)

var db *sql.DB

// InitDB initialize the MySQL database
func InitDB(dsn string) error {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	if _, err = db.Exec(schema); err != nil {
		return err
	}
	return nil
}

// StoreCandle stores the completed candle in the MySQL database
func StoreCandle(symbol string, candle *models.Candle) error {
	log.Printf("StoreCandle: symbol: %s, candle: %v", symbol, candle)
	_, err := db.Exec(`INSERT INTO candlesticks (symbol, timestamp, open, high, low, close, volume)
        VALUES (?, ?, ?, ?, ?, ?, ?)`,
		symbol, candle.Timestamp, candle.Open, candle.High, candle.Low, candle.Close, candle.Volume)
	return err
}

// GetFilterCandles gets the candles data from the MySQL database
func GetFilterCandles(symbol, period string, numberOfCandles int) ([]models.ChartCandle, error) {
	var data []models.ChartCandle

	query := "SELECT symbol,timestamp, open, high, low,close,volume  FROM candlesticks WHERE symbol = ?"
	rows, err := db.Query(query, symbol)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var candle models.ChartCandle
		var startTimeRaw []byte
		if err := rows.Scan(&candle.Symbol, &startTimeRaw, &candle.Open, &candle.High, &candle.Low, &candle.Close, &candle.Volume); err != nil {
			return data, err
		}
		startTime, err := time.Parse(time.RFC3339, string(startTimeRaw))
		if err != nil {
			return data, err
		}

		candle.Timestamp = startTime
		data = append(data, candle)
	}
	if len(data) == 0 {
		return nil, errors.New("no candles found")
	}
	return data, nil
}

// MySQL schema for the candles table
const schema = `
CREATE TABLE IF NOT EXISTS candlesticks (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    symbol VARCHAR(10) NOT NULL,
    timestamp DATETIME NOT NULL,
    open DOUBLE NOT NULL,
    high DOUBLE NOT NULL,
    low DOUBLE NOT NULL,
    close DOUBLE NOT NULL,
    volume DOUBLE NOT NULL,
    UNIQUE KEY unique_candle (symbol, timestamp)
);
`
