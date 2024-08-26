# Build Stage
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o loadgen main.go

# Final Stage
FROM alpine:latest

# Install certificates for secure communication
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/loadgen .

# Expose any ports the app is expected to run on (if applicable, but not needed for loadgen as it's a client)
# No need to expose any ports since loadgen acts as a client
# EXPOSE 50052

# Command to run the executable
CMD ["./loadgen"]
