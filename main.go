package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type AnalyticsData struct {
	Time    time.Time `json:"time"`
	UserID  string    `json:"user_id"`
	Headers struct {
		UserAgent     string `json:"user_agent"`
		Authorization string `json:"authorization"`
	} `json:"headers"`
	Body json.RawMessage `json:"body"`
}

type Config struct {
	Port          int    `json:"port"`
	DBCredentials string `json:"db_credentials"`
	LogLevel      string `json:"log_level"`
}

var (
	analyticsQueue = make(chan AnalyticsData, 100) // Buffer for analytics data
	config         Config
)

func main() {
	loadConfig() // Load configuration from a file or other source

	// Start worker pool for processing analytics data
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ { // Number of worker goroutines
		wg.Add(1)
		go processAnalyticsData(&wg)
	}

	r := mux.NewRouter()
	r.HandleFunc("/analytics", handleAnalytics).Methods("POST")

	http.Handle("/", r)
	address := fmt.Sprintf(":%d", config.Port)
	log.Printf("Listening on %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handleAnalytics(w http.ResponseWriter, r *http.Request) {
	var data AnalyticsData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Send data to the analytics queue (non-blocking)
	select {
	case analyticsQueue <- data:
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"status": "OK"}`))
	default:
		http.Error(w, "Server Busy", http.StatusServiceUnavailable)
	}
}

func processAnalyticsData(wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range analyticsQueue {
		// Simulate data processing and saving to the database
		time.Sleep(time.Second)
		log.Printf("Processed analytics data for user %s\n", data.UserID)
	}
}

func loadConfig() {
	// Load config from file or other source (example below)
	config.Port = 8080
	config.DBCredentials = "your_db_credentials"
	config.LogLevel = "info"
}
