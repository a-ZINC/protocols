package core

import (
	"fmt"
	"log"
	"net"
)

type ServerClient struct {}

// tcp server creation
func (t *ServerClient) Listen(serverType string, port string, responseHandler func(conn net.Conn)) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	log.Printf("%s server listening on %v \n", serverType, listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("failed to accept: %v", err)
		}
		go responseHandler(conn)
	}
}

// client try to connect to server
func (t *ServerClient) Dial(port string, domain string, requestHandler func(conn net.Conn)) error {
	conn, err := net.Dial("tcp", domain + ":" + port)
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()

	requestHandler(conn)
	return nil
}
