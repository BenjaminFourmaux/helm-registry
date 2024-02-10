FROM golang:1.21-bullseye

WORKDIR /app

# Copy source code file and folder
COPY . .

# Install go modules
RUN go mod download

# Build the application
RUN go build -o main .

# Expose the port
EXPOSE 8080
