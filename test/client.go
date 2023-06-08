package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 连接到服务器
	conn, err := net.Dial("tcp", "100.76.246.116:8333")
	if err != nil {
		fmt.Println("无法连接到服务器：", err)
		return
	}
	defer conn.Close()

	// 读取用户输入的命令
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入命令：")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取命令失败：", err)
			break
		}

		// 去除命令中的换行符
		cmd = strings.TrimSpace(cmd)

		// 发送命令给服务器
		_, err = conn.Write([]byte(cmd + "\n"))
		if err != nil {
			fmt.Println("发送命令失败：", err)
			break
		}

		// 接收服务器的响应
		response := make([]byte, 1024) // 假设响应不超过 1024 字节
		n, err := conn.Read(response)
		if err != nil {
			fmt.Println("接收响应失败：", err)
			break
		}

		// 打印服务器的响应
		fmt.Println("服务器响应：", string(response[:n]))
	}
}
