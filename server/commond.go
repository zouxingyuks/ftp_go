package server

import (
	"fmt"
	"ftp_go/server/config"
	"net"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// 处理客户端命令

func processCommand(dialog *WorkSpace, command string, arguments []string) []byte {
	switch command {
	case "AUTH":
		//todo 待实现
		return nil
	case "USER":
		//todo  设置传输类型
		// 用户名命令，返回成功
		return []byte("331 User OK\r\n")
	case "PASS":
		//todo  设置传输类型
		// 密码命令，返回登录成功
		return []byte("230 Login OK\r\n")
	case "TYPE":
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

	case "NOOP":
		// NOOP 命令，空操作，保持连接活跃
		return []byte("200 OK\r\n")
	case "PWD":
		//todo 设置基准路径
		// PWD 命令，返回当前工作目录路径
		//Dir := config.Configs.GetString("Dir.root") + "/" + dialog.Usr + dialog.Dir
		dir := "/" + dialog.Dir
		response := fmt.Sprintf("257 \"%s\" is the current directory.\r\n", dir)
		return []byte(response)
	case "CWD":
		// CWD 命令，切换当前工作目录
		if len(arguments) < 1 {
			return []byte("500 CWD command requires a directory argument.\r\n")
		}
		// 获取客户端传递的目录参数
		fmt.Println("cwd :", arguments[0])
		if arguments[0] == "/" {
			dialog.Dir = ""
			return []byte("250 Directory successfully changed.\r\n")
		}
		dir := dialog.Dir + "/" + arguments[0]
		dialog.Dir = dir
		// 返回成功响应给客户端
		return []byte("250 Directory successfully changed.\r\n")

	case "QUIT":
		// 退出命令，返回成功并关闭连接
		return []byte("221 Goodbye\r\n")
	case "PORT":
		// PORT 命令，客户端指定数据连接地址和端口
		if len(arguments) < 1 {
			return []byte("500 PORT command requires an argument.\r\n")
		}

		// 解析客户端指定的数据连接地址和端口
		address, err := parsePortAddress(arguments[0])
		if err != nil {
			return []byte(fmt.Sprintf("500 Failed to parse PORT argument: %s\r\n", err))
		}
		// 建立数据连接
		dataConn, err := net.Dial("tcp", address)
		if err != nil {
			return []byte(fmt.Sprintf("500 Failed to establish data connection: %s\r\n", err))
		}
		dialog.DataConn = dataConn
		return []byte("200 Data connection established.\r\n")
	case "CDUP":
		// CDUP 命令，将当前工作目录切换到父级目录
		dialog.Dir = filepath.Dir(dialog.Dir)
		// 返回成功响应给客户端
		return []byte("200 OK\r\n")
	case "RETR":

	case "EPRT":
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

	case "LIST":
		// 列出文件列表的命令
		dir := dialog.Dir
		//todo Dir 应该是默认为当前工作目录
		if len(arguments) > 0 {
			dir = arguments[0]
		}
		dir = config.Configs.GetString("Dir.root") + "/" + dialog.Usr + dir
		fmt.Println("list Dir: ", dir)
		cmd := exec.Command("bash", "-c", fmt.Sprintf("ls -l %s | tail -n +2", dir))
		output, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			sendResponse(dialog.CommandConn, []byte("500 List Error\r\n"), dialog.TransferType) // 发送错误响应给客户端
			return nil
		}

		// 发送成功的响应给客户端
		sendResponse(dialog.CommandConn, []byte("200 List OK\r\n"), dialog.TransferType)
		sendResponse(dialog.DataConn, append(output, []byte("\r\n")...), dialog.TransferType) // 发送文件列表数据给客户端
		//todo  不知道对不对
		dialog.DataConn.Close()
		return nil

	default:
		// 未知命令，返回错误
		return []byte("500 Unknown Command\r\n")
	}
	return nil
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

// 辅助函数：将字符串转换为整数
func toInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
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
