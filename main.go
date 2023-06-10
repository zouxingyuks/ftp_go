package main

import (
	"fmt"
	"ftp_go/Connect"
	"ftp_go/config"
	"github.com/sirupsen/logrus"
	"net"
)

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
	go HandleMonitoring()

	// 持续运行
	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受连接失败：", err)
			continue
		}

		// 并行处理客户端连接
		go Connect.HandleClient(conn)
	}
}

// 处理监控
func HandleMonitoring() {
	for {
		// 监控逻辑...
		//fmt.Println("监控中...")
		//
		//time.Sleep(5 * time.Second)
	}
}

func main() {
	Start()
}
