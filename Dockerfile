# Use the official Ubuntu base image
FROM ubuntu:latest

# Install necessary packages including Go, FFMPEG, Python, and PIP
RUN apt-get update && apt-get install -y \
    ffmpeg \
    python3 \
    python3-pip \
    golang-go \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory in the container to /app
WORKDIR /app

# Copy the go.mod and go.sum files first (for better layer caching)
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application; adjust the path according to your project structure
RUN go build -o mucut ./cmd

EXPOSE 8080
# Command to run the executable
CMD ["./mucut"]
