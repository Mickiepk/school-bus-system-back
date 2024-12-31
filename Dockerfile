FROM golang:1.17-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from the official PostgreSQL image
FROM postgres:13-alpine

# Set the Current Working Directory inside the container
WORKDIR /docker-entrypoint-initdb.d

# Copy the SQL script into the container
COPY Servers/v1/db/UserAccount.sql /docker-entrypoint-initdb.d/

# Copy the built Go app from the builder stage
COPY --from=builder /app/main /app/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the PostgreSQL server and the Go app
CMD ["sh", "-c", "docker-entrypoint.sh postgres & /app/main"]