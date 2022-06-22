# syntax=docker/dockerfile:1


FROM golang:1.18.2-alpine3.15

WORKDIR /app

COPY go.mod ./main/
COPY go.sum ./main/

COPY ../zmq_connector/go.mod ./zmq_connector/

WORKDIR /app/main
RUN go mod download
WORKDIR /app
COPY dgraph-http/*.go ./main/

WORKDIR /app/zmq_connector
RUN go mod download
WORKDIR /app
COPY zmq_connector/*.go ./zmq_connector/

WORKDIR /app/main
RUN go build  -o /hStream-dgraph

CMD [ "/hStream-dgraph" ]
