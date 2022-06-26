module solenopsys.org/zmq_dgraph

go 1.17

replace solenopsys.org/zmq_connector v0.0.0 => ./zmq_connector

require solenopsys.org/zmq_connector v0.0.0

require (
	github.com/go-zeromq/goczmq/v4 v4.2.2 // indirect
	github.com/go-zeromq/zmq4 v0.14.1 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
)
