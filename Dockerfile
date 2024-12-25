# Build stage
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests to the container
COPY go.mod go.sum ./

# Download the module dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

# Final stage: Minimal runtime image
FROM scratch

# Set the working directory inside the final container
WORKDIR /

# Copy the statically built binary from the build stage
COPY --from=builder /app/main .

# Command to run the binary
CMD ["/main"]