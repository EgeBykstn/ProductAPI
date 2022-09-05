# syntax=docker/dockerfile:1
FROM golang:1.18 AS builder

WORKDIR /app
COPY . .
CMD go build main.go
ENTRYPOINT ["/app/main"]
#RUN go mod tidy
#RUN go build -o main main.go

#FROM golang:1.18
#WORKDIR /app
#COPY --from=builder /app/main .

#EXPOSE 1323

#ENTRYPOINT ["/app/main"]