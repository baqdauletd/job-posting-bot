# Start from an official Go base image
FROM golang:1.24.2 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go app
RUN go build -o bot

# Final stage: small runtime image
FROM debian:bookworm-slim

# Install certificates for HTTPS
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Create app directory
WORKDIR /app

# Copy built binary and env
COPY --from=builder /app/bot .
COPY .env .

# Run the bot
CMD ["./bot"]
