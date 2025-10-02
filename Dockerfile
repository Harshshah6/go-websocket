# Use the official Go image as a build stage
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go module files and download deps
COPY go.mod ./
RUN go mod tidy

# Copy the source
COPY . .

# Build the Go app
RUN go build -o server main.go

# Runtime stage
FROM debian:bullseye-slim

WORKDIR /app

# Copy the built server and static files
COPY --from=builder /app/server .
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./server"]
