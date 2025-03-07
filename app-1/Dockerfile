# First stage: Build the Go application
FROM golang:1.20 AS builder

WORKDIR /app

# Copy source code
COPY main.go .

# Build the Go app with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# Second stage: Create a minimal runtime image
FROM alpine:latest

WORKDIR /app

# Install necessary dependencies (if needed)
RUN apk --no-cache add ca-certificates

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Ensure execution permission
RUN chmod +x /app/server

# Expose port 8090
EXPOSE 8090

# Run the application
CMD ["/app/server"]
