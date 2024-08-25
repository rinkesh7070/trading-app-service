package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
	"sync"
	"trading-app-service/internal/aggregator"
	"trading-app-service/internal/binance"
	"trading-app-service/internal/db"
	"trading-app-service/internal/server"
)

var (
	DSN, gRPCServerPort, port string
	Symbols                   []string
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DBNAME")
	gRPCServerPort = os.Getenv("GRPC_SERVER_PORT")
	Symbols = strings.Split(os.Getenv("BINANCE_SYMBOLS"), ",")

	// Get the port from environment variable, default to 8080 if not set
	port = os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	DSN = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, host, dbname)
}

func main() {
	// Initialize Aggregator
	agg := aggregator.NewAggregator()

	// Initialize Database
	if err := db.InitDB(DSN); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Printf("Connected to database.")

	// Initialize Binance Client
	client, err := binance.NewBinanceClient(Symbols)
	if err != nil {
		log.Fatalf("Failed to create Binance client: %v", err)
	}
	log.Printf("Connected to Binance.")

	// Initialize gRPC Server
	grpcServer, err := server.StartGRPCServer(agg, gRPCServerPort)
	if err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
	log.Printf("gRPC server started on port %s", gRPCServerPort)

	// Start the HTTP server
	go func() {
		if err := server.StartHTTPServer(port); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()
	log.Printf("HTTP server started on port %s", port)

	// Use a wait group to manage goroutines
	var wg sync.WaitGroup
	wg.Add(1)

	// Concurrently handle message processing
	go func() {
		defer wg.Done()
		server.ProcessMessages(client, grpcServer, agg)
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
