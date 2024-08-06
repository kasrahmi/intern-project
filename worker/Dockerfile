# Build Stage
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app for the correct architecture
RUN GOOS=linux GOARCH=amd64 go build -o worker main.go

# Add this line to verify the binary is created
RUN ls -l /app/worker

# Final Stage
FROM alpine:latest

# Install CA certificates if needed for HTTPS
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/worker .

# Verify that the binary exists in the final image
RUN ls -l /app/worker

# Command to run the executable
CMD ["./worker"]
