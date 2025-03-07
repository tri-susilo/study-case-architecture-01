# Use an official Go image as a builder
FROM golang:1.21 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files, then download dependencies
COPY go.mod go.sum ./ 
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-crud

# Use a minimal base image
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /go-crud /app/go-crud

# Copy the views directory into the container
COPY --from=builder /app/views /app/views

# Ensure the binary has execution permission
RUN chmod +x /app/go-crud

EXPOSE 8080

# Command to run the application
CMD ["/app/go-crud"]
