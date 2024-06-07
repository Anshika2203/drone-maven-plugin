# Use golang:1.22 as the builder stage
FROM golang:1.22 as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies (go.sum may be empty if there are no dependencies)
RUN go mod download

# Copy the source code
COPY . .

# Build the Go executable for amd64 architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /app/drone-maven-plugin ./main.go

# Use openjdk:latest as the base image
FROM maven:3.8.4-jdk-11

# Copy the built Go executable from the builder stage to the final stage
COPY --from=builder /app/drone-maven-plugin /app/

# Set the entrypoint to the Go executable
ENTRYPOINT ["/app/drone-maven-plugin"]