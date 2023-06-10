package Command

import (
	"ftp_go/models"
	"strings"
)

func handleTYPE(dialog *models.WorkSpace, arguments []string) []byte {
	// TYPE 命令，设置传输类型
	if len(arguments) < 1 {
		return []byte("500 TYPE command requires an argument.\r\n")
	}

	// 获取客户端传递的参数
	transferType := strings.ToUpper(arguments[0])

	// 根据客户端传递的参数设置传输类型
	switch transferType {
	case "A", "ASCII":
		// 设置传输类型为 ASCII
		dialog.TransferType = "ASCII"
		return []byte("200 Type set to ASCII.\r\n")

	case "I", "BINARY":
		// 设置传输类型为 binary
		dialog.TransferType = "BINARY"
		return []byte("200 Type set to binary.\r\n")

	default:
		// 无效的传输类型参数
		return []byte("500 Unknown transfer type.\r\n")
	}
}
