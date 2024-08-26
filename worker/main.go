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
    pb "worker/proto"
)

type server struct {
    pb.UnimplementedMyServiceServer
    id string
}

func (s *server) ProcessRequest(ctx context.Context, req *pb.MyRequest) (*pb.MyResponse, error) {
    startTime := time.Now()
    duration := time.Duration(req.GetDurationSeconds()) * time.Millisecond
    endTime := startTime.Add(duration)

    for time.Now().Before(endTime) {}

    latency := time.Since(startTime).Milliseconds()
    response := &pb.MyResponse{
        Message:           "Processed " + req.GetName(),
        EndToEndLatencyMs: latency,
        WorkerId:          s.id,
    }
    return response, nil
}

func main() {
    port := flag.String("port", "8080", "The server port")
    if envPort, exists := os.LookupEnv("PORT"); exists {
        *port = envPort
    }
    id := *port
    flag.Parse()

    lis, err := net.Listen("tcp", ":"+*port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterMyServiceServer(s, &server{id: id})
    reflection.Register(s)

    log.Printf("Starting gRPC server on port %s with ID %s...", *port, id)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
