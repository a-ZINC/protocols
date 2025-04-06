package http

import (
	"bufio"
	"custom-protocols/core"
	"fmt"
	"log"
	"net"
	"strings"
)

var methodSet = map[string]int {
	"GET": 1,
	"POST": 2,
	"PUT": 3,
	"DELETE": 4,
}

type Http struct {
	server core.ServerClient
	port   string
}

func New(port string) *Http {
	return &Http{
		server: core.ServerClient{},
		port:   port,
	}
}

func (t *Http) StartServer() {
	err := t.server.Listen("chttp", t.port, t.handleConnection)
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}

func (t *Http) handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("Http connection from %s", conn.RemoteAddr())
	reader := bufio.NewReader(conn)

	// method checking
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("unable to read data, %v", err)
	}
	if err := t.isValidMethodHeader(str); err != nil {
		log.Printf("unable to process request, %v", err)
		return
	}
	log.Printf("request method is valid %v", str)

	// Headers parsing


	

}

func (t *Http) isValidMethod(method string) bool {
	_, exist := methodSet[method]
	return exist
}

func (t *Http) isValidPath(path string) bool {
	return strings.HasPrefix(path, "/")
}

func (t *Http) isValidMethodHeader(method string) error {
	methods := strings.Split(strings.TrimSpace(method), " ")
	// method[0] as chttp , method[1] as path, method[2] as GET,POST etc
	if len(methods) != 3 {
		return fmt.Errorf("invalid method header")
	}
	if (methods[0] != "chttp") {
		return fmt.Errorf("%v method is not supported", methods[0])
	}
	if !t.isValidMethod(methods[1]) {
		return fmt.Errorf("%v method is not supported", methods[1])
	}
	if !t.isValidPath(methods[2]) {
		return fmt.Errorf("%v path is not supported", methods[2])
	}
	return nil
}