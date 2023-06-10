package Command

import (
	"fmt"
	"ftp_go/Response"
	"ftp_go/models"
	"strings"
)

type HandlerFunc func(dialog *models.WorkSpace, command string, arguments []string) []byte

func Handle(dialog *models.WorkSpace, handlerFunc HandlerFunc) {
	for {
		// 读取客户端命令
		cmdStr, err := dialog.Reader.ReadString('\n')
		fmt.Println(cmdStr)
		if err != nil {
			fmt.Println("读取命令失败：", err)
			break
		}
		// 去除命令中的换行符和回车符
		cmd := strings.TrimRight(cmdStr, "\r\n")
		// 解析命令和参数
		tokens := strings.Split(cmd, " ")
		command := strings.ToUpper(tokens[0])
		arguments := tokens[1:]
		// 处理客户端命令
		fmt.Println(command)
		response := handlerFunc(dialog, command, arguments)
		Response.Send(dialog.CommandConn, response, dialog.TransferType)
		if command == "QUIT" {
			break
		}
	}
}

// 处理客户端命令

func Process(dialog *models.WorkSpace, command string, arguments []string) []byte {
	//登录状态检测
	if !dialog.Status {
		switch command {
		case "USER":
			// 用户名命令，返回成功
			return HandleUSER(dialog, arguments)
		case "PASS":
			// 密码命令，返回成功
			return HandlePASS(dialog, arguments)
		case "QUIT":
			// 退出命令，返回成功并关闭连接
			return []byte("221 Goodbye\r\n")
		default:
			return []byte("530 Please login with USER and PASS.\r\n")
		}
	}

	switch command {
	case "AUTH":
		//todo 待实现
		return nil
	case "USER":
		// 用户名命令，返回成功
		return HandleUSER(dialog, arguments)
	case "PASS":
		// 密码命令，返回成功
		return HandlePASS(dialog, arguments)
	case "TYPE":
		// TYPE 命令，设置传输类型
		return HandleTYPE(dialog, arguments)
	case "STOR":
		// STOR 命令，上传文件
		return HandleSTOR(dialog, arguments)
	case "NOOP":
		// NOOP 命令，空操作，保持连接活跃
		return []byte("200 OK\r\n")
	case "PWD":
		// PWD 命令，返回当前工作目录路径
		return HandlePWD(dialog)
	case "CWD":
		// CWD 命令，切换当前工作目录
		return HandleCWD(dialog, arguments)

	case "QUIT":
		// 退出命令，返回成功并关闭连接
		return []byte("221 Goodbye\r\n")
	case "PORT":
		// PORT 命令，客户端指定数据连接地址和端口
		return HandlePORT(dialog, arguments)
	case "CDUP":
		// CDUP 命令，将当前工作目录切换到父级目录
		return HandleCDUP(dialog)
	case "RETR":
		// RETR 命令，下载文件
		return HandleRETR(dialog, arguments[0])

	case "EPRT":
		// EPRT 命令，客户端指定数据连接地址和端口
		return HandleEPRT(dialog, arguments)

	case "LIST":
		return HandleList(dialog, arguments)
	case "DELE":
		return HandleDELE(dialog, arguments)
	case "RNFR":
		return HandleRNFR(dialog, arguments)
	case "RNTO":
		return HandleRNTO(dialog, arguments)
	case "MKD":
		return HandleMKD(dialog, arguments)
	default:
		// 未知命令，返回错误
		return []byte("500 Unknown Command\r\n")
	}

	return nil
}
