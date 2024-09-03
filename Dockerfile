FROM docker.io/golang:1.23.0 AS builder

WORKDIR /app

COPY . .

RUN make

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ip .

EXPOSE 8080

CMD ["/app/ip"]
