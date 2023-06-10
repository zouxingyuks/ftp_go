package Command

import (
	"ftp_go/Response"
	"ftp_go/models"
	"os"
	"path/filepath"
	"syscall"
)

func handleSTOR(dialog *models.WorkSpace, arguments []string) []byte {
	if len(arguments) < 1 {
		return []byte("501 缺少文件名。\r\n")
	}

	// 获取要上传的文件名
	filename := filepath.Join(dialog.Dir, arguments[0])

	// 验证权限逻辑
	if !checkPermissions(dialog, filename) {
		dialog.Logs.Warnln("尝试在没有权限的情况下存储文件。")
		return []byte("550 权限被拒绝。\r\n")
	}

	Response.Send(dialog.CommandConn, []byte("150 打开数据连接。\r\n"), dialog.TransferType)

	// 从数据连接中读取数据
	dataConn := dialog.DataConn

	// 打开文件以写入数据
	f, err := os.OpenFile(filepath.Join(dialog.BasicDir, dialog.Dir, arguments[0]), syscall.O_WRONLY|syscall.O_CREAT|syscall.O_TRUNC, os.ModePerm)
	if err != nil {
		dialog.Logs.Errorln("打开文件时出错：", err)
		return nil
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			dialog.Logs.Errorln("关闭文件时出错：", err)
		}
	}(f)

	// 读取数据并写入文件
	data := make([]byte, 1024)
	for {
		n, err := dataConn.Read(data)
		if err != nil {
			break
		}
		write, err := f.Write(data[:n])
		if err != nil {
			dialog.Logs.Errorln("写入数据连接时出错：", err)
			return nil
		}
		if write != n {
			dialog.Logs.Errorln("写入数据连接时出错：", err)
			return nil
		}
	}

	// 发送传输完成响应
	Response.Send(dialog.CommandConn, []byte("226 传输完成。\r\n"), dialog.TransferType)

	return nil
}
