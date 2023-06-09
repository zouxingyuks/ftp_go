package Command

import (
	"ftp_go/models"
	"os"
	"path/filepath"
)

func HandleCWD(dialog *models.WorkSpace, arguments []string) []byte {
	if len(arguments) < 1 {
		return []byte("500 CWD command requires a directory argument.\r\n")
	}
	var newDir string
	if filepath.IsAbs(arguments[0]) {
		// 如果是绝对路径，直接使用
		newDir = arguments[0]
	} else {
		// 如果是相对路径，与当前工作目录进行合并
		newDir = filepath.Join(dialog.Dir, arguments[0])
	}
	// 检查新的工作目录是否存在
	if _, err := os.Stat(filepath.Join(dialog.BasicDir, newDir)); os.IsNotExist(err) {
		return []byte("550 No such directory." + newDir + "\r\n")
	}

	// 更新工作目录
	dialog.Dir = newDir

	// 返回成功响应给客户端
	return []byte("250 Directory successfully changed.\r\n")
}
