package main

import (
    "context"
    "flag"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "log"
    "net"
    "time"
    pb "worker/proto"  // Adjust to match your module path
)

type server struct {
    pb.UnimplementedMyServiceServer
    id string  // Unique identifier for the Worker instance
}

func (s *server) ProcessRequest(ctx context.Context, req *pb.MyRequest) (*pb.MyResponse, error) {
    startTime := time.Now()

    // Simulate CPU-bound work
    duration := time.Duration(req.GetDurationSeconds()) * time.Second
    endTime := startTime.Add(duration)
    for time.Now().Before(endTime) {
        // Busy-wait to simulate CPU-bound task
    }

    latency := time.Since(startTime).Milliseconds()

    response := &pb.MyResponse{
        Message:                "Processed " + req.GetName(),
        EndToEndLatencyMs:      latency,
        WorkerId:               s.id,  // Add the Worker ID to the response
    }
    return response, nil
}

func main() {
    port := flag.String("port", "50052", "The server port")  // Default to 50052 or any available port
    id := flag.String("id", "worker-1", "Unique identifier for this worker")
    flag.Parse()

    lis, err := net.Listen("tcp", ":"+*port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterMyServiceServer(s, &server{id: *id})
    reflection.Register(s)
    log.Printf("Starting gRPC server on port %s with ID %s...", *port, *id)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
