# syntax=docker/dockerfile:1
FROM product-api:controller

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o product-api

EXPOSE 5432

CMD ["./ProductAPI"]