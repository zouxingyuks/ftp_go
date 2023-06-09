package Command

import (
	"fmt"
	"ftp_go/Response"
	"ftp_go/models"
	"io/ioutil"
	"os"
	"path/filepath"
)

func HandleRETR(dialog *models.WorkSpace, filename string) []byte {
	// 构建完整的文件路径
	path := filepath.Join(dialog.BasicDir, dialog.Dir, filename)

	// 检查文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 文件不存在，返回错误
		return []byte(fmt.Sprintf("550 %s: No such file or directory\r\n", filename))
	}

	// 读取文件内容
	data, err := ioutil.ReadFile(path)
	if err != nil {
		// 读取文件失败，返回错误
		return []byte(fmt.Sprintf("550 Failed to open file: %s\r\n", err))
	}

	// 发送开始传输的响应
	Response.Send(dialog.CommandConn, []byte("150 Opening data connection\r\n"), dialog.TransferType)

	// 发送文件内容
	Response.Send(dialog.DataConn, append(data, []byte("\r\n")...), dialog.TransferType)
	dialog.DataConn.Close()

	// 发送传输完成的响应
	Response.Send(dialog.CommandConn, []byte("226 Transfer complete\r\n"), dialog.TransferType)

	// 返回nil表示没有错误
	return nil
}
