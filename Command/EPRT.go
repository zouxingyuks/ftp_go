package Command

import (
	"fmt"
	"ftp_go/models"
	"net"
	"strings"
)

func handleEPRT(dialog *models.WorkSpace, arguments []string) []byte {
	// EPRT 命令，客户端指定数据连接地址和端口（IPv4或IPv6）
	if len(arguments) < 1 {
		return []byte("500 EPRT command requires an argument.\r\n")
	}

	// 解析客户端指定的数据连接地址和端口
	address, err := parseEprtAddress(arguments[0])
	if err != nil {
		return []byte(fmt.Sprintf("500 Failed to parse EPRT argument: %s\r\n", err))
	}
	// 建立数据连接
	dataConn, err := net.Dial("tcp", address)
	if err != nil {
		return []byte(fmt.Sprintf("500 Failed to establish data connection: %s\r\n", err))
	}
	dialog.DataConn = dataConn
	return []byte("200 Data connection established.\r\n")
}

// 解析 EPRT 命令参数，获取数据连接地址和端口
func parseEprtAddress(arg string) (string, error) {
	parts := strings.Split(arg, "|")
	if len(parts) != 5 {
		return "", fmt.Errorf("Invalid EPRT argument")
	}

	ip := parts[2]
	port := toInt(parts[3]) // 注意，这里没有位移操作，因为EPRT命令参数的格式是|net_proto|net_addr|tcp_port|

	address := fmt.Sprintf("[%s]:%d", ip, port) // 对于IPv6地址，需要用[]括起来
	return address, nil
}
