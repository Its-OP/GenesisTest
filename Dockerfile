# Start from the latest Golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY ./application ./application
COPY ./domain ./domain
COPY ./infrastructure ./infrastructure
COPY ./web ./web

COPY ./docs ./docs
COPY ./startup.go .

# Build the Go app
RUN go build -o main startup.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
