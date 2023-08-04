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

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/server/main .

# Copy the .env file and firebaseconfig.json from the build stage
COPY --from=builder /app/.env .
COPY --from=builder /app/firebaseconfig.json .

# Expose port 5050 for the API service
EXPOSE 5050

# Set the working directory to /app/server/ and the entry point of the container
WORKDIR /app/server/

# Set the entry point of the container
CMD ["./main"]
