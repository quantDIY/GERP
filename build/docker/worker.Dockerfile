# build/docker/worker.Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /worker ./cmd/worker

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /worker .
ENTRYPOINT ["./worker"]
