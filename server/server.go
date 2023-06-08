package server

import (
	"bufio"
	"fmt"
	"ftp_go/server/config"
	"github.com/sirupsen/logrus"
	"net"
	"strings"
)

type CommandHandlerFunc func(dialog *WorkSpace, command string, arguments []string) []byte

var serverLogs *logrus.Entry

func init() {
	config.LoadDefaultConfig()
	config.InitLog()
	config.ParseConfig()
	// 创建服务日志
	serverLogs = config.NewLog("server")
}

func Start() {
	fmt.Println("启动FTP服务器...")

	// 监听地址
	listenAddress := config.Configs.GetString("host.ip") + ":" + config.Configs.GetString("host.port")

	// 创建监听器
	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		fmt.Println("监听失败：", err)
		serverLogs.Errorln("监听失败：", err)
		return
	}
	defer listener.Close()

	fmt.Println("监听地址：", listenAddress)

	// 处理监控
	go handleMonitoring()

	// 持续运行
	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受连接失败：", err)
			continue
		}

		// 并行处理客户端连接
		go handleClient(conn)
	}
}

// 处理监控
func handleMonitoring() {
	for {
		// 监控逻辑...
		//fmt.Println("监控中...")
		//
		//time.Sleep(5 * time.Second)
	}
}

// 处理客户端连接
func handleClient(conn net.Conn) {
	defer conn.Close()
	dialog := &WorkSpace{
		CommandConn:  conn,
		Reader:       bufio.NewReader(conn),
		Dir:          "",
		TransferType: "",
	}
	// 验证登录
	err := authenticate(dialog)
	if err != nil {
		fmt.Println("登录验证失败：", err)
		return
	}
	handleConn(dialog, processCommand)
	// 处理客户端请求

}

// 验证登录
func authenticate(dialog *WorkSpace) error {
	// 向客户端发送登录提示
	_, err := dialog.CommandConn.Write([]byte("220 Please enter your username:\r\n"))
	if err != nil {
		return err
	}
	// 读取客户端的用户名
	username, err := dialog.Reader.ReadString('\n')
	if err != nil {
		return err
	}
	username = strings.TrimSpace(username)

	// 向客户端发送进一步验证的提示
	_, err = dialog.CommandConn.Write([]byte("331 Please enter your password:\r\n"))
	if err != nil {
		return err
	}

	// 读取客户端的密码
	password, err := dialog.Reader.ReadString('\n')
	if err != nil {
		return err
	}
	password = strings.TrimSpace(password)
	username = strings.Replace(username, "USER ", "", -1)
	password = strings.Replace(password, "PASS ", "", -1)
	// 执行登录验证逻辑
	if !checkCredentials(username, password, dialog) {
		// 登录验证失败，向客户端发送错误消息并关闭连接
		_, err = dialog.CommandConn.Write([]byte("530 Login incorrect.\r\n"))
		if err != nil {
			return err
		}
		return fmt.Errorf("登录失败：用户名或密码错误")
	}

	// 登录验证成功，向客户端发送登录成功的消息
	_, err = dialog.CommandConn.Write([]byte("230 Login successful.\r\n"))
	if err != nil {
		return err
	}

	return nil
}

// 检查用户名和密码
func checkCredentials(username, password string, dialog *WorkSpace) bool {
	// 在这里进行实际的用户名和密码验证逻辑
	// 例如，从数据库或配置文件中验证用户凭据

	// 假设用户名为 "admin"，密码为 "password"
	// 这里只是一个示例，你需要根据实际情况进行修改
	if username == "admin" && password == "password" {
		//dialog.Usr = username
		dialog.Usr = "anubis"
		return true
	}

	return false
}

func handleConn(dialog *WorkSpace, handlerFunc CommandHandlerFunc) {
	for {
		// 读取客户端命令
		cmdBytes, err := dialog.Reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取命令失败：", err)
			break
		}
		// 去除命令中的换行符和回车符
		cmd := strings.TrimRight(cmdBytes, "\r\n")
		// 解析命令和参数
		tokens := strings.Split(cmd, " ")
		command := strings.ToUpper(tokens[0])
		arguments := tokens[1:]
		// 处理客户端命令
		response := handlerFunc(dialog, command, arguments)
		sendResponse(dialog.CommandConn, response, dialog.TransferType)

	}
}

// 发送响应给客户端
func sendResponse(conn net.Conn, response []byte, transferType string) {
	switch transferType {
	case "BINARY":
		{
			// 如果传输类型为 BINARY，则将响应数据作为二进制流发送
			_, err := conn.Write(response)
			if err != nil {
				fmt.Println("发送响应失败：", err)
			}
		}
	case "ASCII":
		{
			// 如果传输类型为 ASCII，则将响应数据作为字符串发送
			_, err := fmt.Fprint(conn, string(response))
			if err != nil {
				fmt.Println("发送响应失败：", err)
			}
		}
	default:
		{
			_, err := conn.Write(response)
			if err != nil {
				fmt.Println("发送响应失败：", err)
			}
		}
	}
}
