# Stage 1: Build the Go binary
FROM golang:1.23.0 as builder

WORKDIR /trading-app-service
COPY . .

RUN go mod download
RUN go mod tidy
RUN go build -o trading-app-service ./cmd/main.go
RUN chmod +x trading-app-service

# Stage 2: Create the minimal image
FROM alpine:3.18

WORKDIR /trading-app-service

# Copy the binary from the builder stage
COPY --from=builder /trading-app-service/trading-app-service /trading-app-service/

# Copy the schema.sql file if needed
COPY --from=builder /trading-app-service/internal/db/schema.sql /schema.sql

EXPOSE 50051

CMD ["./trading-app-service"]
