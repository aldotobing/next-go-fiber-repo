# Use the official GoLang Docker image with Go 1.19 as the base image
FROM golang:1.19-alpine

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

# Expose port 5050 for the API service
EXPOSE 5000

# Set the entry point of the container
CMD ["./main"]
