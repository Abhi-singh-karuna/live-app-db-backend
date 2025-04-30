# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy SQL schema file
COPY --from=builder /app/internal/infrastructure/database/init.sql ./internal/infrastructure/database/

# Create a non-root user
RUN adduser -D -g '' appuser
USER appuser

# Expose port 8080
EXPOSE 8080

# Set environment variables
ENV TURSO_DATABASE_URL="libsql://test-db-devofgolang.aws-us-east-1.turso.io"
ENV TURSO_AUTH_TOKEN="eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3NDYwMTc2MjksImlkIjoiZDI2NTI5MjMtYTFlOC00NmM2LWEzNWUtN2Y4ODlmNjA4M2ZhIiwicmlkIjoiZTE0NjUyNTUtOTJjYS00YWVkLThmN2EtMDYyNTA2OWFiYTIzIn0.JVnvEisuq2JypCHNDXKyQJC0aNCTqX-yPocc8Ve8r9yub9Jw1LDtDGnAnl_ZEFjiK9kiMAC6864ygGRRBTldCw"

# Run the application
CMD ["./main"]
