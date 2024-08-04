package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)

type Metrics struct {
    Ack          string  `json:"ack"`
    E2ELatencyMs float64 `json:"e2e_latency_ms"`
}

func cpuSpin(duration time.Duration) {
    end := time.Now().Add(duration)
    for time.Now().Before(end) {
    }
}

func workerHandler(w http.ResponseWriter, r *http.Request) {
    startTime := time.Now()

    duration := 100 * time.Millisecond
    if d, err := time.ParseDuration(r.URL.Query().Get("duration")); err == nil {
        duration = d
    }

    cpuSpin(duration)

    e2eLatency := time.Since(startTime).Seconds() * 1000

    metrics := Metrics{
        Ack:          "Request processed",
        E2ELatencyMs: e2eLatency,
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(metrics); err != nil {
        log.Printf("Failed to encode response: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/worker", workerHandler)
    port := ":8080"
    fmt.Printf("Worker function is running on port %s\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
