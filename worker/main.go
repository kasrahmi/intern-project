package main

import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "log"
    "net"
    "time"
    pb "worker/proto"  // Adjust to match your module path
)

type server struct {
    pb.UnimplementedMyServiceServer
}

func (s *server) ProcessRequest(ctx context.Context, req *pb.Request) (*pb.Response, error) {
    startTime := time.Now()
    
    // Simulate CPU-bound work
    duration := time.Duration(req.GetDurationSeconds()) * time.Second
    endTime := startTime.Add(duration)
    for time.Now().Before(endTime) {
        // Busy-wait to simulate CPU-bound task
    }

    latency := time.Since(startTime).Milliseconds()

    response := &pb.Response{
        Message:                "Processed " + req.GetName(),
        EndToEndLatencyMs: latency,
    }
    return response, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterMyServiceServer(s, &server{})
    reflection.Register(s)
    log.Println("Starting gRPC server on port 50051...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
