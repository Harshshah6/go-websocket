# Build stage
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod and go.sum and download deps
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire project
COPY . .

# Build the Go server from cmd/server
RUN go build -o server ./cmd/server

# Runtime stage
FROM debian:bullseye-slim

WORKDIR /app

# Copy the built binary
COPY --from=builder /app/server .

# Copy static files
COPY --from=builder /app/static ./static

# Expose default port 8080
EXPOSE 8080

# Run the server
CMD ["./server"]
