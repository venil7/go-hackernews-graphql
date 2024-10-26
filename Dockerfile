# Stage 1: Build the Go server
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go server binary
RUN go build -o server .

# Stage 2: Run the Go server
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Expose the port the server will listen on
EXPOSE 8080

# Run the Go server
CMD ["./server"]
