package Command

import (
	"ftp_go/models"
	"path/filepath"
)

func handleCDUP(dialog *models.WorkSpace) []byte {
	if dialog.Dir == "/" {
		// 已经在根目录，无法再上升，返回错误
		return []byte("550 Can't change directory\r\n")
	}

	// 使用 filepath.Dir 函数更容易地获取父目录
	dir := filepath.Dir(dialog.Dir)
	// 修正根目录路径，确保它是 "/"
	if dir == "." {
		dir = "/"
	}

	dialog.Dir = dir
	// 返回成功响应给客户端
	return []byte("250 Directory successfully changed.\r\n")
}
