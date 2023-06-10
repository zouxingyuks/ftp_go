package Command

import (
	"fmt"
	"ftp_go/models"
	"os"
	"path/filepath"
)

func handleDELE(dialog *models.WorkSpace, arguments []string) []byte {
	// 获取要删除的文件路径
	Path := filepath.Join(dialog.BasicDir, dialog.Dir, arguments[0])
	// 检查新的工作目录是否存在
	if _, err := os.Stat(filepath.Join(Path)); os.IsNotExist(err) {
		return []byte("550 No such directory." + Path + "\r\n")
	}
	fmt.Println("delete:", Path)
	// 执行删除文件的操作，你可以根据你的实际需求来实现
	err := deleteFile(Path)
	if err != nil {
		// 删除文件失败，返回错误消息
		return []byte(fmt.Sprintf("550 Failed to delete file: %s\r\n", err))
	}

	// 文件删除成功，返回成功消息
	return []byte("250 File deleted successfully\r\n")
}

func deleteFile(filePath string) error {
	// 在这里实现删除文件的逻辑
	// 你可以使用 os 包来删除文件
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
