package Command

import (
	"fmt"
	"ftp_go/models"
)

func handlePWD(dialog *models.WorkSpace) []byte {
	//todo 设置基准路径
	// PWD 命令，返回当前工作目录路径
	dir := dialog.Dir
	response := fmt.Sprintf("257 \"%s\" is the current directory.\r\n", dir)
	return []byte(response)
}
