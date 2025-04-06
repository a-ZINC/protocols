package main

import (
	"custom-protocols/protocols/http"
	"custom-protocols/protocols/tcp"
	"flag"
	"fmt"
)

func main() {
	types := flag.String("type", "server", "type of Viewer")
	protocol := flag.String("protocol", "tcp", "protocol to use")
	port := flag.String("port", "8080", "port to use")
	flag.Parse()

	switch *types {
	case "server":
		switch *protocol {
		case "tcp":
			tcp.New(*port).StartServer()
		case "chttp":
			http.New(*port).StartServer()
		default:
			fmt.Printf("protocol not supported")
		}
	case "client":
		fmt.Printf("client")
	}
}
