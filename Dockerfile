# Build stage
FROM golang:1.23-alpine AS builder

# Install git
RUN apk add --no-cache git

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go

# Run stage
FROM alpine:latest

WORKDIR /root/

# Install ca-certificates
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port
EXPOSE 8080

# Command to run
CMD ["./main"]
