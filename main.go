package main

import (
	"bytes"
	"fmt"
	zmq_connector "github.com/solenopsys/sc-bl-zmq-connector"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func main() {
	template := zmq_connector.HsTemplate{Pf: processingFunction()}
	template.Init()
}
