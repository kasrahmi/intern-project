package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    "time"
    pb "load-generator/proto"  // Adjust to match your module path
)

const (
    address         = "localhost:50052"  // Target Worker on port 50052
    requestsPerSec  = 5  // Adjust the rate as needed
    durationSec     = 5  // Duration for each request's CPU-bound task
)

func sendRequest(client pb.MyServiceClient) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)  // Increased timeout
    defer cancel()

    req := &pb.Request{
        Name:             "LoadTest",
        DurationSeconds:  durationSec,
    }
    resp, err := client.ProcessRequest(ctx, req)
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Response from port %s: %s, Latency: %d ms", address, resp.GetMessage(), resp.GetEndToEndLatencyMs())
}

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    client := pb.NewMyServiceClient(conn)

    ticker := time.NewTicker(time.Second / time.Duration(requestsPerSec))
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            go sendRequest(client)
        }
    }
}
