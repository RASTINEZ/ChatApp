# Start from the official Golang image to build the app
FROM golang:1.24.2-alpine as builder

WORKDIR /app

# Copy go.mod and go.sum, then download dependencies (this is cached unless changed)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your code
COPY . .

# Build the Go app (replace 'main.go' with your main file if named differently)
RUN go build -o main .

# Start a fresh, minimal image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your Go app runs on (change if your app uses a different port)
EXPOSE 5000

# Command to run the binary
CMD ["./main"]