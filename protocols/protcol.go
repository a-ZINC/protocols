package protocols

import "net"

type Protocol interface {
	StartServer(port string, handleConnection func(conn *net.Conn)) error
	handleConnection(conn *net.Conn)
}
