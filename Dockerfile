FROM golang:1.23.5-alpine AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./api

FROM debian:latest

WORKDIR /bin
COPY --from=builder /build/api ./api

CMD ["/bin/api"]
