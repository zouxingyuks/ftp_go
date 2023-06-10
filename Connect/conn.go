package Connect

import (
	"bufio"
	"ftp_go/Command"
	"ftp_go/config"
	"ftp_go/models"
	"net"
)

// 处理客户端连接
func HandleClient(conn net.Conn) {
	defer conn.Close()
	dialog := &models.WorkSpace{
		CommandConn:  conn,
		Reader:       bufio.NewReader(conn),
		Dir:          "/",
		BasicDir:     config.Configs.GetString("dir.root"),
		TransferType: "",
		Status:       false,
	}
	dialog.CommandConn.Write([]byte("220 Please enter your username:\r\n"))
	Command.Handle(dialog)
	// 处理客户端请求

}
