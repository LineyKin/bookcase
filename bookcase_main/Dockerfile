# Этап сборки
FROM golang:1.23.3-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o myapp

# Этап запуска
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /usr/src/app/myapp .

CMD ["./myapp"]