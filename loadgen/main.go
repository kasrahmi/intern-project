package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    pb "load-generator/proto" // Ensure this matches the generated code's package
)

func main() {
    conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    c := pb.NewMyServiceClient(conn)

    resp, err := c.MyMethod(context.Background(), &pb.MyRequest{Name: "World"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Response: %s", resp.GetMessage())
}
