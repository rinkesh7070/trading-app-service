package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"trading-app-service/internal/chart"
	"trading-app-service/internal/db"
)

// Home Handler for the homepage
func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/templates/home.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Could not execute template", http.StatusInternalServerError)
	}
}

// GetCandleStickChart Handler for the candlestick chart page
func GetCandleStickChart(w http.ResponseWriter, r *http.Request) {
	log.Printf("Enter in GetCandleStickChart")
	query := r.URL.Query()

	period := "1d"
	symbol := "BTCUSDT"

	periodQuery := query.Get("period")
	if periodQuery != "" {
		period = periodQuery
	}

	symbolQuery := query.Get("symbol")
	if symbolQuery != "" {
		symbol = symbolQuery
	}

	fmt.Printf("Retrieving candles for %s in period %s\n", symbol, period)
	chartCandles, err := db.GetFilterCandles(symbol, period, 100)
	if err != nil {
		fmt.Printf("Failed to retrieve candles for %s in period %s: %v\n", symbol, period, err)
		http.Error(w, "No result found in DB.", http.StatusOK)
		return
	}

	csPage := chart.NewCandleStickPage()
	csPage.AddCandleStickChart(symbol, chartCandles)
	csPage.AddMarkPoint("buy", "2020/12/15 21:00", 22000.0)
	csPage.AddMarkPoint("sell", "2021/02/07 21:00", 46000.0)
	csPage.Render(w)
}
