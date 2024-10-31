# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.23.2-alpine AS builder
WORKDIR /app
# Copy all files from the current directory to the working directory in the container
COPY . .
# Build the Go application with optimizations
RUN go build -ldflags "-s -w" -o main .

# Run stage
FROM alpine
# Install necessary packages
RUN apk add --no-cache tzdata
# Set the timezone
ENV TZ=Asia/Bangkok

WORKDIR /app

# Create a system group and user for running the application
RUN addgroup --system --gid 1001 golanggroup
RUN adduser --system --uid 1001 golang

# Copy the built application and environment file from the builder stage
COPY --from=builder --chown=golang:golanggroup /app/main .
COPY --chown=golang:golanggroup ./app.env .

# Create a directory for images, set ownership and permissions
RUN mkdir -p /app/image && \
    chown golang:golanggroup /app/image && \
    chmod 755 /app/image

# Switch to the non-root user
USER golang

# Expose the application port
EXPOSE 8080
# Command to run the application
CMD ["/app/main"]