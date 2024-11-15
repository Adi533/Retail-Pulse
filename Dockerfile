# Use the official Go image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first to leverage Docker cache for dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application files
COPY . .

# Build the Go binary
RUN go build -o main .

# Expose the port the app will run on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
