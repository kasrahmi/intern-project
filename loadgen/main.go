package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "log"
    "math/rand"
    "os"
    "time"

    pb "load-generator/proto" // Adjust to match your module path
)

var workerAddresses = []string{
    "localhost:50052",
    "localhost:50053",
}

// sendRequest sends a request to a worker and logs the results.
func sendRequest(client pb.MyServiceClient, logFile *os.File, rps int, distribution string, duration int) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
    defer cancel()

    req := &pb.MyRequest{
        Name:            "LoadTest",
        DurationSeconds: int32(duration),
    }
    startTime := time.Now()
    resp, err := client.ProcessRequest(ctx, req)
    endTime := time.Now()

    if err != nil {
        log.Printf("could not process: %v", err)
        logFile.WriteString(fmt.Sprintf("Time: %v, RPS: %d, Distribution: %s, Duration: %d ms, Request: %v, Error: %v\n",
            endTime.Format(time.RFC3339),
            rps,
            distribution,
            duration,
            req.GetName(),
            err))
    } else {
        latency := endTime.Sub(startTime).Milliseconds()
        logFile.WriteString(fmt.Sprintf("Time: %v, RPS: %d, Distribution: %s, Duration: %d ms, Request: %v, Response: %v, Latency: %d ms, Worker ID: %s\n",
            endTime.Format(time.RFC3339),
            rps,
            distribution,
            duration,
            req.GetName(),
            resp.GetMessage(),
            latency,
            resp.GetWorkerId()))
    }
}

// generatePoissonInterval generates an interval based on a Poisson distribution.
func generatePoissonInterval(lambda float64) time.Duration {
    // Use exponential distribution for Poisson process
    interval := rand.ExpFloat64() / lambda
    return time.Duration(interval * float64(time.Second))
}

func main() {
    var clients []pb.MyServiceClient
    for _, address := range workerAddresses {
        conn, err := grpc.Dial(address, grpc.WithInsecure())
        if err != nil {
            log.Fatalf("did not connect to %s: %v", address, err)
        }
        clients = append(clients, pb.NewMyServiceClient(conn))
    }

    // Create a log file
    logFile, err := os.Create("loadgen_log.txt")
    if err != nil {
        log.Fatalf("failed to create log file: %v", err)
    }
    defer logFile.Close()

    // Test with different RPS and CPU-spin duration settings
    for rps := 5; rps <= 50; rps += 5 {
        for duration := 100; duration <= 1000; duration += 100 {
            for _, distribution := range []string{"Uniform", "Poisson"} {
                startTime := time.Now()
                logFile.WriteString(fmt.Sprintf("Starting test with %d RPS, %s distribution, %d ms duration\n", rps, distribution, duration))
                totalDuration := 10 * time.Second
                endTime := startTime.Add(totalDuration)

                ticker := time.NewTicker(time.Second / time.Duration(rps))
                defer ticker.Stop()

                for time.Now().Before(endTime) {
                    select {
                    case <-ticker.C:
                        client := clients[rand.Intn(len(clients))] // Randomly select a worker
                        go sendRequest(client, logFile, rps, distribution, duration)
                    default:
                        if distribution == "Poisson" {
                            time.Sleep(generatePoissonInterval(float64(rps)))
                            client := clients[rand.Intn(len(clients))] // Randomly select a worker
                            go sendRequest(client, logFile, rps, distribution, duration)
                        }
                    }
                }

                log.Printf("Test with %d RPS, %s distribution, %d ms duration completed\n", rps, distribution, duration)
            }
        }
    }
}
