package Command

import (
	"fmt"
	"ftp_go/models"
	"os"
	"path/filepath"
)

func handleMKD(dialog *models.WorkSpace, arguments []string) []byte {
	if len(arguments) < 1 {
		return []byte("501 Missing directory name.\r\n")
	}

	// 获取要创建的目录路径
	Path := filepath.Join(dialog.Dir, arguments[0])

	// 验证权限逻辑
	if !checkPermissions(dialog, Path) {
		//低权限用户危险操作报警
		dialog.Logs.Warnln("Try to create " + Path + " but permission denied.")
		return []byte("550 Permission denied.\r\n")
	}

	// 检查目录是否已存在
	if _, err := os.Stat(filepath.Join(dialog.BasicDir, Path)); !os.IsNotExist(err) {
		dialog.Logs.Errorln("Try to create " + Path + " but it already exists.")
		return []byte("550 Directory already exists.\r\n")
	}
	// 执行创建目录操作
	err := createDir(filepath.Join(dialog.BasicDir, Path))
	if err != nil {
		// 创建目录失败，返回错误消息
		dialog.Logs.Errorln("Try to create "+Path+" but failed. err:", err)
		return []byte(fmt.Sprintf("550 Failed to create directory: %s\r\n", err))
	}

	// 目录创建成功，返回成功消息
	dialog.Logs.Infoln("Directory " + Path + " created.")
	return []byte("257 Directory created\r\n")
}

func createDir(directoryPath string) error {
	err := os.Mkdir(directoryPath, 0777) // 设置目录的权限为777，根据需要进行调整
	return err
}
