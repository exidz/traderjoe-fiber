# Use the official Go image as the base image
FROM golang:1.20-alpine
# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o myapi main.go

# Expose the port your API listens on
EXPOSE 3000

# Command to run the application
CMD ["./myapi"]
