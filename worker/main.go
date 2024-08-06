package main

import (
    "context"
    "flag"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "log"
    "net"
    "os"
    "time"
    pb "worker/proto" // Adjust to match your module path
)

type server struct {
    pb.UnimplementedMyServiceServer
    id string // Unique identifier for the Worker instance
}

// ProcessRequest handles incoming gRPC requests
func (s *server) ProcessRequest(ctx context.Context, req *pb.MyRequest) (*pb.MyResponse, error) {
    startTime := time.Now()

    // Simulate CPU-bound work
    duration := time.Duration(req.GetDurationSeconds()) * time.Second
    endTime := startTime.Add(duration)
    for time.Now().Before(endTime) {
        // Busy-wait to simulate CPU-bound task
        // Consider using a more efficient way to simulate work if needed
    }

    latency := time.Since(startTime).Milliseconds()

    // Create the response with message, latency, and worker ID
    response := &pb.MyResponse{
        Message:           "Processed " + req.GetName(),
        EndToEndLatencyMs: latency,
        WorkerId:          s.id,
    }
    return response, nil
}

func main() {
    // Retrieve port from environment variable or use the default
    port := flag.String("port", "50052", "The server port")
    if envPort, exists := os.LookupEnv("PORT"); exists {
        *port = envPort
    }

    // Unique identifier for this worker
    id := port
    flag.Parse()

    // Listen on the specified TCP port
    lis, err := net.Listen("tcp", ":"+*port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // Create a new gRPC server
    s := grpc.NewServer()
    pb.RegisterMyServiceServer(s, &server{id: *id})

    // Register reflection service on gRPC server for debugging and testing
    reflection.Register(s)

    log.Printf("Starting gRPC server on port %s with ID %s...", *port, *id)

    // Serve gRPC server
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
