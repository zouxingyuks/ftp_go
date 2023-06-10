package Command

import (
	"fmt"
	"ftp_go/models"
	"os"
	"path/filepath"
)

func handleRMD(dialog *models.WorkSpace, arguments []string) []byte {
	if len(arguments) < 1 {
		return []byte("501 Missing directory name.\r\n")
	}
	// 获取要删除的目录路径
	Path := filepath.Join(dialog.Dir, arguments[0])
	// 验证权限逻辑
	if !checkPermissions(dialog, Path) {
		//低权限用户危险操作报警
		dialog.Logs.Warnln("Try to delete " + Path + " but permission denied.")
		return []byte("550 Permission denied.\r\n")
	}

	// 检查目录是否存在，如果不存在，则不能移除
	if _, err := os.Stat(filepath.Join(dialog.BasicDir, Path)); os.IsNotExist(err) {
		dialog.Logs.Errorln("Try to delete " + Path + " but not exist.")
		return []byte("550 Directory not found.\r\n")
	}

	// 执行删除目录操作
	err := deleteDir(filepath.Join(dialog.BasicDir, Path))
	if err != nil {
		// 删除目录失败，返回错误消息
		dialog.Logs.Errorln("Try to delete "+Path+" but failed. err:", err)
		return []byte(fmt.Sprintf("550 Failed to remove directory: %s\r\n", err))
	}

	// 目录删除成功，返回成功消息
	dialog.Logs.Infoln("Directory " + Path + " removed.")
	return []byte("250 Directory removed\r\n")
}
func deleteDir(Path string) error {
	err := os.Remove(Path)
	return err
}
