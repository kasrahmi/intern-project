package main

import (
    "context"
    "fmt"
    "log"
    "math/rand"
    "os"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    pb "load-generator/proto"
)

var workerAddresses = []string{
    "worker.default.svc.cluster.local:8080", // Replace with your Knative service name
}

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
        log.Printf("Time: %v, RPS: %d, Distribution: %s, Duration: %d ms, Request: %v, Response: %v, Latency: %d ms, Worker ID: %s\n",
            endTime.Format(time.RFC3339),
            rps,
            distribution,
            duration,
            req.GetName(),
            resp.GetMessage(),
            latency,
            resp.GetWorkerId())
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

func main() {
    var clients []pb.MyServiceClient
    for _, address := range workerAddresses {
        conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
        if err != nil {
            log.Fatalf("did not connect to %s: %v", address, err)
        }
        clients = append(clients, pb.NewMyServiceClient(conn))
    }

    logFile, err := os.Create("loadgen_log.txt")
    if err != nil {
        log.Fatalf("failed to create log file: %v", err)
    }
    defer logFile.Close()

    rps := 10
    startTime := time.Now()
    distribution := "Uniform"
    duration := 500

    ticker := time.NewTicker(time.Second / time.Duration(rps))
    defer ticker.Stop()

    totalDuration := 10 * time.Second
    endTime := startTime.Add(totalDuration)
    logFile.WriteString(fmt.Sprintf("Starting test with %d RPS, %s distribution, %d ms duration\n", rps, distribution, duration))

    for time.Now().Before(endTime) {
        select {
        case <-ticker.C:
            client := clients[rand.Intn(len(clients))]
            go sendRequest(client, logFile, rps, distribution, duration)
        }
    }

    log.Printf("Test with %d RPS, %s distribution, %d ms duration completed\n", rps, distribution, duration)
}
