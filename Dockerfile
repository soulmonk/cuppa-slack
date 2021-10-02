# Builder image
FROM golang:1.17.1-alpine AS builder

WORKDIR /build
COPY . /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o application .

# Exec image
FROM alpine:latest

COPY --from=builder /build/application /application

ENTRYPOINT ["/application"]