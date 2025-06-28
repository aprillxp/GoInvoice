# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.21-alpine AS build
WORKDIR /app

# Install git for Go modules that need it
RUN apk add --no-cache git

# Copy go.mod & go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy entire project
COPY . .

# Build binary
RUN go build -o main ./cmd/api

# Final stage (minimal)
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=build /app/main .

# Expose port
EXPOSE 8080

# Run binary
CMD ["./main"]
