package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    url := "http://localhost:8080/worker"
    rps := 10
    interval := time.Second / time.Duration(rps)

    ticker := time.NewTicker(interval)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            go sendRequest(url)
        }
    }
}

func sendRequest(url string) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Failed to send request: %v\n", err)
        return
    }
    defer resp.Body.Close()

    fmt.Printf("Request sent. Response status: %s\n", resp.Status)
}
