# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN GOOS=linux GOARCH=amd64 go build -o /log-hello

FROM busybox

COPY --from=builder /log-hello /log-hello

ENTRYPOINT [ "/log-hello" ]