package server

import (
	"bufio"
	"net"
)

type workSpace struct {
	commondConn net.Conn
	dataConn    net.Conn
	reader      *bufio.Reader
	usr         string
	//相对路径
	dir          string
	transferType string
}
