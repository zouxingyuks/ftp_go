package Command

import (
	"ftp_go/models"
	"os"
	"path/filepath"
)

func handleRNTO(dialog *models.WorkSpace, arguments []string) []byte {
	// 检查是否已经执行过 RNFR 命令
	if dialog.RNFR == "" {
		return []byte("503 Bad sequence of commands.\r\n")
	}

	// 检查是否有权限重命名文件
	if !checkPermissions(dialog, dialog.RNFR) {
		return []byte("550 Permission denied.\r\n")
	}

	// 重命名文件
	err := os.Rename(filepath.Join(dialog.BasicDir, dialog.RNFR), filepath.Join(dialog.BasicDir, dialog.Dir, arguments[0]))
	if err != nil {
		return []byte("550 Failed to rename file.\r\n")
	}

	// 清空 RNFR
	dialog.RNFR = ""

	// 返回成功响应给客户端
	return []byte("250 File renamed successfully.\r\n")
}
