// package main

// import (
//     "context"
//     "fmt"
//     "google.golang.org/grpc"
//     "log"
//     "math"
//     "math/rand"
//     "os"
//     "time"
//     pb "load-generator/proto" // Adjust to match your module path
// )

// const (
//     durationSec = 5
// )

// var workerAddresses = []string{
//     "localhost:50052",
//     "localhost:50053",
// }

// // sendRequest sends a single request to the worker and logs the result.
// func sendRequest(client pb.MyServiceClient, logFile *os.File) {
//     ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
//     defer cancel()

//     req := &pb.MyRequest{
//         Name:            "LoadTest",
//         DurationSeconds: durationSec,
//     }
//     startTime := time.Now()
//     resp, err := client.ProcessRequest(ctx, req)
//     endTime := time.Now()

//     if err != nil {
//         log.Printf("could not greet: %v", err)
//         logFile.WriteString(fmt.Sprintf("Time: %v, Error: %v\n", endTime.Format(time.RFC3339), err))
//     } else {
//         latency := endTime.Sub(startTime).Milliseconds()
//         logFile.WriteString(fmt.Sprintf("Time: %v, Request: %v, Response: %v, Latency: %d ms, Worker ID: %s\n",
//             endTime.Format(time.RFC3339),
//             req.GetName(),
//             resp.GetMessage(),
//             latency,
//             resp.GetWorkerId()))
//     }
// }

// // runTest runs the load test with the specified RPS and distribution.
// func runTest(rps int, distribution string, logFile *os.File) {
//     startTime := time.Now()

//     logFile.WriteString(fmt.Sprintf("Starting test with %d requests per second, distribution: %s\n", rps, distribution))
//     duration := 10 * time.Second
//     endTime := startTime.Add(duration)
//     ticker := time.NewTicker(time.Second / time.Duration(rps))
//     defer ticker.Stop()

//     // Function to calculate the next delay for Poisson distribution
//     nextPoissonInterval := func(lambda float64) time.Duration {
//         return time.Duration(-math.Log(1.0-rand.Float64()) / lambda * float64(time.Second))
//     }

//     lambda := float64(rps)

//     for time.Now().Before(endTime) {
//         switch distribution {
//         case "uniform":
//             <-ticker.C // Wait for the ticker

//         case "poisson":
//             // Calculate the next interval using Poisson distribution
//             time.Sleep(nextPoissonInterval(lambda))
//         }

//         // Randomly select a worker and send the request
//         client := clients[rand.Intn(len(clients))]
//         go sendRequest(client, logFile)
//     }
//     log.Printf("Test with %d requests per second and %s distribution completed\n", rps, distribution)
// }

// var clients []pb.MyServiceClient

// func main() {
//     for _, address := range workerAddresses {
//         conn, err := grpc.Dial(address, grpc.WithInsecure())
//         if err != nil {
//             log.Fatalf("did not connect to %s: %v", address, err)
//         }
//         clients = append(clients, pb.NewMyServiceClient(conn))
//     }

//     // Create a log file
//     logFile, err := os.Create("loadgen_log_distribution.txt")
//     if err != nil {
//         log.Fatalf("failed to create log file: %v", err)
//     }
//     defer logFile.Close()

//     // Test both uniform and Poisson distributions
//     distributions := []string{"uniform", "poisson"}
//     for _, distribution := range distributions {
//         for rps := 5; rps <= 50; rps += 5 {
//             runTest(rps, distribution, logFile)
//         }
//     }
// }