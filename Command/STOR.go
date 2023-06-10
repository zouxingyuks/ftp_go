package Command

import (
	"ftp_go/Response"
	"ftp_go/models"
	"os"
	"path/filepath"
	"syscall"
)

func HandleSTOR(dialog *models.WorkSpace, arguments []string) []byte {
	Response.Send(dialog.CommandConn, []byte("150 Opening data connection.\r\n"), dialog.TransferType)
	// 从数据连接中读取数据
	dataConn := dialog.DataConn
	// 读取数据
	data := make([]byte, 1024)
	f, _ := os.OpenFile(filepath.Join(dialog.BasicDir, dialog.Dir, arguments[0]), syscall.O_WRONLY|syscall.O_CREAT|syscall.O_TRUNC, os.ModePerm)
	//todo 文件创造的错误处理
	for {
		n, err := dataConn.Read(data)
		if err != nil {
			break
		}
		f.Write(data[:n])
	}
	//存储数据
	Response.Send(dialog.CommandConn, []byte("226 Transfer complete.\r\n"), dialog.TransferType)
	return nil
}
