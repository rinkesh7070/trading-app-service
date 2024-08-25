package server

import (
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
	"trading-app-service/internal/aggregator"
	"trading-app-service/internal/binance"
	"trading-app-service/internal/db"
	"trading-app-service/internal/models"
	pb "trading-app-service/internal/proto"
)

type Server struct {
	pb.UnimplementedTradingServiceServer
	aggregator  *aggregator.Aggregator
	subscribers map[string][]chan *pb.Candle
	mu          sync.Mutex
}

func NewServer(agg *aggregator.Aggregator) *Server {
	return &Server{
		aggregator:  agg,
		subscribers: make(map[string][]chan *pb.Candle),
	}
}

func (s *Server) StreamCandles(req *pb.CandleRequest, stream pb.TradingService_StreamCandlesServer) error {
	ch := make(chan *pb.Candle)
	s.mu.Lock()
	s.subscribers[req.Symbol] = append(s.subscribers[req.Symbol], ch)
	s.mu.Unlock()

	for candle := range ch {
		if err := stream.Send(candle); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) BroadcastCandle(symbol string, candle *models.Candle) {
	s.mu.Lock()
	subscribers := s.subscribers[symbol]
	s.mu.Unlock()

	protoCandle := &pb.Candle{
		Symbol:    symbol,
		Timestamp: candle.Timestamp.Unix(),
		Open:      candle.Open,
		High:      candle.High,
		Low:       candle.Low,
		Close:     candle.Close,
		Volume:    candle.Volume,
	}

	for _, ch := range subscribers {
		ch <- protoCandle
	}
}

func ProcessMessages(client *binance.Client, grpcServer *Server, agg *aggregator.Aggregator) {
	for {
		message, err := client.ReadMessages()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			time.Sleep(1 * time.Second) // Sleep to avoid tight loop on errors
			continue
		}

		symbol := message["s"].(string)

		priceStr := message["p"].(string)
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Printf("Error converting price to float64: %v", err)
			continue
		}

		volumeStr := message["q"].(string)
		volume, err := strconv.ParseFloat(volumeStr, 64)
		if err != nil {
			log.Printf("Error converting volume to float64: %v", err)
			continue
		}

		// Process Tick Data
		completeCandle := agg.ProcessTick(symbol, price, volume)
		if completeCandle != nil {
			if err := db.StoreCandle(symbol, completeCandle); err != nil {
				log.Printf("Error storing completeCandle data: %v", err)
				continue
			}
			grpcServer.BroadcastCandle(symbol, completeCandle)
		}
	}
}

func StartGRPCServer(agg *aggregator.Aggregator, address string) (*Server, error) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer()
	server := NewServer(agg)
	pb.RegisterTradingServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	return server, nil
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Route to serve the homepage
	r.HandleFunc("/", Home).Methods("GET")

	// Route to serve the candlestick chart page
	r.HandleFunc("/candlestickchart", GetCandleStickChart).Methods("GET")

	// Serve static files like CSS and JS from the public directory
	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return r
}

func StartHTTPServer(addr string) error {
	router := NewRouter()
	return http.ListenAndServe(addr, router)
}
