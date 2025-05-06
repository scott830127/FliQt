FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o fliqt ./cmd/fliqt/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/fliqt ./fliqt
COPY ./deploy/config/config.local.toml ./config.toml
EXPOSE 8080
ENTRYPOINT ["./fliqt", "-b", "./config.toml"]
