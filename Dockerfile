# Use the official GoLang Docker image with Go 1.19 as the base image for build
FROM golang:1.19-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the .env file and firebaseconfig.json to the working directory
COPY .env .env
COPY firebaseconfig.json firebaseconfig.json

# Copy the rest of the application source code to the working directory
COPY . .

# Change the working directory to the /server subdirectory
WORKDIR /app/server

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Start a new stage from a minimal Debian base image
FROM debian:stable-slim

# Install CA certificates for HTTPS connections and set the timezone to Indonesia/Jakarta
RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Create a non-root user and group
RUN addgroup --system appgroup && adduser --system --ingroup appgroup appuser

# Set the working directory
WORKDIR /app/server

# Copy the pre-built binary file from the previous stage to /app/server/
COPY --from=builder /app/server/main .

# Copy the .env file and firebaseconfig.json from the build stage to /app/
COPY --from=builder /app/.env /app/
COPY --from=builder /app/firebaseconfig.json /app/

# Set file permissions for configuration files
RUN chown appuser:appgroup /app/.env /app/firebaseconfig.json && \
    chmod 600 /app/.env /app/firebaseconfig.json

# Switch to non-root user
USER appuser

# Expose port 5000 for the API service
EXPOSE 5000

# Define the entry point of the container
CMD ["./main"]
