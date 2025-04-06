package protocols

import "net"

type Protocol interface {
	StartServer()
	handleConnection(conn *net.Conn)
}
