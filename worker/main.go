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
    port string
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
        Message:                "Processed " + req.GetName() + " on port " + s.port,
        EndToEndLatencyMs: latency,
    }
    return response, nil
}

func main() {
    port := flag.String("port", "50052", "The server port")  // Use port 50052 or any other port
    flag.Parse()

    lis, err := net.Listen("tcp", ":"+*port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterMyServiceServer(s, &server{port: *port})
    reflection.Register(s)
    log.Printf("Starting gRPC server on port %s...", *port)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
