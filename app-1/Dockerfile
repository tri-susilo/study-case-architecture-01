# First stage: Build the Go application
FROM golang:1.20 AS builder

WORKDIR /app

# Copy source code
COPY main.go .

# Compile the Go app
RUN go build -o server main.go

# Second stage: Create a minimal runtime image
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Expose port 8090
EXPOSE 8090

# Run the application
CMD ["/app/server"]
