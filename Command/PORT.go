package Command

import (
	"fmt"
	"ftp_go/models"
	"net"
	"strings"
)

func HandlePORT(dialog *models.WorkSpace, arguments []string) []byte {
	if len(arguments) < 1 {
		return []byte("500 PORT command requires an argument.\r\n")
	}

	// 解析客户端指定的数据连接地址和端口
	address, err := parsePortAddress(arguments[0])
	if err != nil {
		Logs.Warnln("Failed to parse PORT argument: ", err)
		return []byte(fmt.Sprintf("500 Failed to parse PORT argument: %s\r\n", err))
	}
	// 建立数据连接
	dataConn, err := net.Dial("tcp", address)
	if err != nil {
		Logs.Warnln("Failed to establish data connection: ", err)
		return []byte(fmt.Sprintf("500 Failed to establish data connection: %s\r\n", err))
	}
	dialog.DataConn = dataConn
	return []byte("200 Data connection established.\r\n")
}

// 解析 PORT 命令参数，获取数据连接地址和端口
func parsePortAddress(arg string) (string, error) {
	parts := strings.Split(arg, ",")
	if len(parts) != 6 {
		return "", fmt.Errorf("Invalid PORT argument")
	}

	ip := strings.Join(parts[:4], ".")
	port := (toInt(parts[4]) << 8) + toInt(parts[5])
	address := fmt.Sprintf("%s:%d", ip, port)
	return address, nil
}
