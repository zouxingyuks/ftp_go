package Response

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net"
)

// 发送响应给客户端
func Send(conn net.Conn, response []byte, transferType string) {
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
			// 如果传输类型为 UTF-8，则将响应数据作为字符串发送
			transformer := transform.NewReader(bytes.NewReader(response), unicode.UTF8.NewEncoder())
			// 读取转换后的结果
			encoded, err := ioutil.ReadAll(transformer)
			if err != nil {
				fmt.Println("编码失败：", err)
				return
			}
			// 输出编码后的字符串
			_, err = conn.Write(encoded)
			if err != nil {
				fmt.Println("发送响应失败：", err)
			}
		}
	}
}
