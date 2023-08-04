# Use the official GoLang Docker image with Go 1.19 as the base image for build
FROM golang:1.19-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the .env file to the working directory
COPY .env ./

# Copy the rest of the application source code to the working directory
COPY . .

# Change the working directory to the /server subdirectory
WORKDIR /app/server

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /app/

# Copy the pre-built binary file from the previous stage to /app/server/
COPY --from=builder /app/server/main /app/server/

# Copy the .env file and firebaseconfig.json from the build stage to /app/
COPY --from=builder /app/.env /app/
COPY --from=builder /app/firebaseconfig.json /app/

# Expose port 5050 for the API service
EXPOSE 5000

# Set the working directory to /app/server/ and the entry point of the container
WORKDIR /app/server/

CMD ["./main"]
