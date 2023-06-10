package Command

import (
	"ftp_go/models"
	"os"
	"path/filepath"
)

func handleRNFR(dialog *models.WorkSpace, arguments []string) []byte {
	if len(arguments) < 1 {
		return []byte("500 RNFR command requires a file argument.\r\n")
	}
	Path := filepath.Join(dialog.Dir, arguments[0])
	// 检查新的工作目录是否存在
	if _, err := os.Stat(filepath.Join(dialog.BasicDir, Path)); os.IsNotExist(err) {
		return []byte("550 " + Path + "doesn't exist" + "\r\n")
	}
	// 检查是否有权限重命名文件
	if !checkPermissions(dialog, Path) {
		return []byte("550 Permission denied.\r\n")
	}
	// 修改RNFR
	dialog.RNFR = Path
	// 返回成功响应给客户端
	return []byte("350 Ready for RNTO.\r\n")
}
