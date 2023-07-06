# Use the official GoLang Docker image with Go 1.19 as the base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the .env file and firebaseconfig.json to the working directory
COPY .env ./
COPY firebaseconfig.json ./

# Copy the rest of the application source code to the working directory
COPY . .

# Change the working directory to the /server subdirectory
WORKDIR /app/server

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

<<<<<<< HEAD
# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/server/main .

# Copy the .env file and firebaseconfig.json from the build stage
COPY --from=builder /app/.env .
COPY --from=builder /app/firebaseconfig.json .

=======
>>>>>>> 52910eb (dockerfile)
# Expose port 5050 for the API service
EXPOSE 5050
<<<<<<< HEAD
=======

<<<<<<< HEAD
# Set the working directory to /app/server/ and the entry point of the container
WORKDIR /app/server/
>>>>>>> f2130ff (Docker and pipeline test)

=======
>>>>>>> 52910eb (dockerfile)
# Set the entry point of the container
CMD ["./main"]
