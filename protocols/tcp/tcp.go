package tcp

import (
	"custom-protocols/core"
	"log"
	"net"
)

type Tcp struct {
	server core.ServerClient
	port   string
}

func New(port string) *Tcp {
	 return &Tcp{
		 server: core.ServerClient{},
		 port: port,
	 }
}
func (t *Tcp) StartServer() {
	err := t.server.Listen(t.port, t.handleConnection)
	if err != nil {
		log.Fatalf("failed to start TCP server: %v", err)
	}
}

func (t *Tcp) Dial(port string, domain string) error {
	return t.server.Dial(port, domain, t.handleConnection)
}

func (t *Tcp) handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("connection from %s", conn.RemoteAddr())
}
