package models

import (
	"bufio"
	"net"
)

type WorkSpace struct {
	CommandConn net.Conn
	DataConn    net.Conn
	Reader      *bufio.Reader
	Usr         string
	Status      bool
	RNFR        string
	//相对路径
	Dir          string
	BasicDir     string
	TransferType string
}
