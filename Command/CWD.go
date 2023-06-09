package Command

import (
	"fmt"
	"ftp_go/models"
	"strings"
)

func HandleCWD(dialog *models.WorkSpace, arguments []string) []byte {
	if len(arguments) < 1 {
		return []byte("500 CWD command requires a directory argument.\r\n")
	}
	if arguments[0] == "/" {
		dialog.Dir = "/"
		return []byte("250 Directory successfully changed.\r\n")
	}
	args := strings.Split("/"+arguments[0], "/")
	dir := dialog.Dir
	for _, str := range args {
		if str == "" {
			continue
		}
		dir = dir + str + "/"
	}
	// 获取客户端传递的目录参数
	fmt.Println("dir :", dir)

	dialog.Dir = dir
	// 返回成功响应给客户端
	return []byte("250 Directory successfully changed.\r\n")
}
