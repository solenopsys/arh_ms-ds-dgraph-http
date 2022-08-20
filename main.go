package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"solenopsys.org/zmq_connector"
	"time"
)

type FuncParam struct {
	uri         string
	contentType string
}

func processingFunction() func(message []byte, functionId uint8) []byte {
	println("START CLIENT")

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	var functions = make(map[uint8]*FuncParam)
	functions[1] = &FuncParam{uri: "/query?timeout=20s", contentType: "application/dql"}
	functions[2] = &FuncParam{uri: "/mutate?commitNow=true", contentType: "application/rdf"}

	host := os.Getenv("dgraph.Host")
	port := os.Getenv("dgraph.Port")
	path := "http://" + host + ":" + port

	return func(message []byte, functionId uint8) []byte {
		fmt.Println("")
		fmt.Println(string(message))
		conf := functions[functionId]
		r := bytes.NewReader(message)

		resp, err := client.Post(path+conf.uri, conf.contentType, r)
		if err != nil {
			log.Print(err)
			return []byte("ERROR_CONNECT")
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
			return []byte("ERROR_READ")
		}

		defer resp.Body.Close()
		fmt.Println("")
		fmt.Println(string(body))
		return body
	}
}

func StreamProcessor(
	stream *zmq_connector.StreamConfig,
	cancel context.CancelFunc,
) {

	f := processingFunction()

	for {
		messageWr := <-stream.Input
		resBytes := f(messageWr.Body, messageWr.Function)
		stream.Output <- &zmq_connector.HsMassage{0, messageWr.Function, resBytes}

		//todo cancel
	}

}

func main() {
	socketUrl := os.Getenv("zmq.SocketUrl")

	streams := &zmq_connector.StreamsHolder{
		Streams:        make(map[uint32]*zmq_connector.StreamConfig),
		Input:          make(chan *zmq_connector.SocketMassage, 256),
		Output:         make(chan *zmq_connector.SocketMassage, 256),
		MessageHandler: StreamProcessor,
	}

	z := &zmq_connector.HsSever{
		SocketUrl: socketUrl,
		Streams:   streams,
	}

	z.StartServer()
}
