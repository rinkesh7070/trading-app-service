package binance

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strings"
)

const baseURL = "wss://stream.binance.com:9443/ws/"

type Client struct {
	conn *websocket.Conn
}

// NewBinanceClient connects to the WebSocket with all required symbols.
func NewBinanceClient(symbols []string) (*Client, error) {
	// Construct the full URL with multiple streams
	streams := make([]string, len(symbols))
	for i, symbol := range symbols {
		streams[i] = fmt.Sprintf("%s@aggTrade", strings.ToLower(symbol))
	}

	fullURL := fmt.Sprintf("%s%s", baseURL, strings.Join(streams, "/"))

	conn, _, err := websocket.DefaultDialer.Dial(fullURL, nil)
	if err != nil {
		log.Printf("BinanceClient dial error: %v", err)
		return nil, err
	}

	client := &Client{conn: conn}

	return client, nil
}

func (b *Client) ReadMessages() (map[string]interface{}, error) {
	var message map[string]interface{}
	err := b.conn.ReadJSON(&message)
	return message, err
}
