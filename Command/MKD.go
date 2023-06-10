package Command

import (
	"fmt"
	"ftp_go/models"
	"os"
	"path/filepath"
)

func HandleMKD(dialog *models.WorkSpace, arguments []string) []byte {
	if len(arguments) < 1 {
		return []byte("501 Missing directory name.\r\n")
	}

	// 获取要创建的目录路径
	Path := filepath.Join(dialog.Dir, arguments[0])

	// 验证权限逻辑
	if !checkPermissions(dialog, Path) {
		return []byte("550 Permission denied.\r\n")
	}

	// 检查目录是否已存在
	if _, err := os.Stat(filepath.Join(dialog.BasicDir, Path)); !os.IsNotExist(err) {
		return []byte("550 Directory already exists.\r\n")
	}
	// 执行创建目录操作
	err := createDirectory(filepath.Join(dialog.BasicDir, Path))
	if err != nil {
		// 创建目录失败，返回错误消息
		return []byte(fmt.Sprintf("550 Failed to create directory: %s\r\n", err))
	}

	// 目录创建成功，返回成功消息
	return []byte("257 Directory created\r\n")
}

func createDirectory(directoryPath string) error {
	// 在这里实现创建目录的逻辑
	// 可以使用操作系统或存储系统的API来创建目录

	err := os.Mkdir(directoryPath, 0777) // 设置目录的权限为777，根据需要进行调整
	if err != nil {
		return err
	}

	return nil
}
