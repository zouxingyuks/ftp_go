package server

import (
	"bufio"
	"net"
)

type WorkSpace struct {
	CommandConn net.Conn
	DataConn    net.Conn
	Reader      *bufio.Reader
	Usr         string
	//相对路径
	Dir          string
	TransferType string
}
