FROM golang:1.21-bullseye

WORKDIR /app

# Copy go mod file
COPY go.mod .

# Install go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Expose the port
EXPOSE 8080
