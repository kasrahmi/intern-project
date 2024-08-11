package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "log"
    "math/rand"
    "os"
    "time"
    pb "load-generator/proto"  // Adjust to match your module path
)

const (
    requestsPerSec = 5
    durationSec    = 5
)

var workerAddresses = []string{
    "localhost:50052",
    "localhost:50053",
}

func sendRequest(client pb.MyServiceClient, logFile *os.File) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
    defer cancel()

    req := &pb.MyRequest{
        Name:             "LoadTest",
        DurationSeconds:  durationSec,
    }
    startTime := time.Now()
    resp, err := client.ProcessRequest(ctx, req)
    endTime := time.Now()

    if err != nil {
        log.Printf("could not greet: %v", err)
        logFile.WriteString(fmt.Sprintf("Time: %v, Error: %v\n", endTime.Format(time.RFC3339), err))
    } else {
        latency := endTime.Sub(startTime).Milliseconds()
        log.Printf("Response: %s, Latency: %d ms, Worker ID: %s", resp.GetMessage(), resp.GetEndToEndLatencyMs(), resp.GetWorkerId())
        logFile.WriteString(fmt.Sprintf("Time: %v, Request: %v, Response: %v, Latency: %d ms, Worker ID: %s\n",
            endTime.Format(time.RFC3339),
            req.GetName(),
            resp.GetMessage(),
            latency,
            resp.GetWorkerId()))
    }
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

    ticker := time.NewTicker(time.Second / time.Duration(requestsPerSec))
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            client := clients[rand.Intn(len(clients))]  // Randomly select a worker
            go sendRequest(client, logFile)
        }
    }
}
