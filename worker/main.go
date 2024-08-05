package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    "net"
    pb "worker/proto" // Import path should match the generated code's package
)

type server struct {
    pb.UnimplementedMyServiceServer // This embeds the default server implementation
}

func (s *server) MyMethod(ctx context.Context, req *pb.MyRequest) (*pb.MyResponse, error) {
    return &pb.MyResponse{Message: "Hello " + req.GetName()}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50052") // Ensure this port matches the client's connection port
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterMyServiceServer(s, &server{}) // Register the MyService server
    log.Println("Starting gRPC server on port 50052...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
