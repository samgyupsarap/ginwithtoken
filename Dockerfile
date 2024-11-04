# Step 1: Build the Go application
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ginapi .

# Step 2: Create a minimal runtime image
FROM alpine:latest

# Install any required libraries
RUN apk add --no-cache ca-certificates

# Set the working directory in the runtime container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/ginapi .

# Expose the port your Gin application runs on
EXPOSE 8080

# Command to run the application
CMD ["./ginapi"]
