package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    "math/rand"
    "time"
    pb "load-generator/proto"  // Adjust to match your module path
)

const (
    requestsPerSec = 5
    durationSec    = 5
)

var workerAddresses = []string{
    "localhost:50052",
    "localhost:50053",  // Add additional worker addresses as needed
}

func sendRequest(client pb.MyServiceClient) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
    defer cancel()

    req := &pb.MyRequest{
        Name:             "LoadTest",
        DurationSeconds:  durationSec,
    }
    resp, err := client.ProcessRequest(ctx, req)
    if err != nil {
        log.Printf("could not greet: %v", err)
    } else {
        log.Printf("Response: %s, Latency: %d ms, Worker ID: %s", resp.GetMessage(), resp.GetEndToEndLatencyMs(), resp.GetWorkerId())
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

    ticker := time.NewTicker(time.Second / time.Duration(requestsPerSec))
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            client := clients[rand.Intn(len(clients))]  // Randomly select a worker
            go sendRequest(client)
        }
    }
}
