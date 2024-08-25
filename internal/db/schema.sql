-- Create the database if it doesn't exist
CREATE DATABASE IF NOT EXISTS trading_data;

-- Use the database
USE trading_data;

-- Create the table to store OHLC candlestick data
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

-- Indexes to speed up queries
CREATE INDEX idx_symbol ON candlesticks(symbol);
CREATE INDEX idx_timestamp ON candlesticks(timestamp);
