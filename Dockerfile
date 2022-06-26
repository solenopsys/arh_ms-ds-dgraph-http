# syntax=docker/dockerfile:1

FROM golang:1.18.2-alpine3.15

WORKDIR /app/main/zmq_connector

COPY zmq_connector/*.go ./
COPY zmq_connector/go.mod ./
RUN go mod download

WORKDIR /app/main
COPY go.mod ./
COPY go.sum ./
COPY *.go ./
RUN go mod download

RUN go build  -o /go-binary

CMD [ "/go-binary" ]

