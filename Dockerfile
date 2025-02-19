# Use an official Go image as the base image
FROM golang:1.18 as builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o /app/surge-pricing ./src

# Use a minimal image for the final build
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built application from the builder
COPY --from=builder /app/surge-pricing .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./surge-pricing"] 