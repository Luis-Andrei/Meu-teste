# Build stage
FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bank-service

# Production stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bank-service .
COPY --from=builder /app/config ./config

EXPOSE 8081

CMD ["./bank-service"] 