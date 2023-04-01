# Use the official Golang image as a base
FROM golang:1.20.2-alpine3.17

# Set the working directory
WORKDIR /app

# Copy the necessary files into the container
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go application
RUN go build -o app .

# Expose the port that the application will run on
EXPOSE 9091

# Set the command to run when the container starts
CMD ["./app"]
