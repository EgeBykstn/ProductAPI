# syntax=docker/dockerfile:1
FROM golang:1.18 AS builder

WORKDIR /app

COPY . .
RUN go build main.go

ENTRYPOINT ["./main"]
