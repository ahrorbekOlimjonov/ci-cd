FROM golang:1.22.5-alpine3.19 AS builder

RUN mkdir app
COPY . /app

WORKDIR /app

RUN go build -o main cmd/main.go

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app .

CMD ["/app/main"]
 