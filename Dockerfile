# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the dependencies
RUN go mod download

# Copy the Go source files into the container
COPY . .

# Compile the application
RUN go build -o app ./cmd/api

# Expose the port on which the application will run
EXPOSE 7777

# Set the default command to run when starting the container
CMD ["./app"]
