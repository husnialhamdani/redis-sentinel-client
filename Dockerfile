# Syntax=docker/dockerfile:1.4

# Build stage
FROM --platform=$BUILDPLATFORM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and application files
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application for the target platform
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o main .

# Run stage
FROM --platform=$TARGETPLATFORM alpine:latest

# Set working directory and copy the built binary
WORKDIR /root/
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]