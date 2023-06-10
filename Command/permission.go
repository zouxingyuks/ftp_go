package Command

import "ftp_go/models"

func checkPermissions(dialog *models.WorkSpace, path string) bool {

	//此处的path 是ftp 中的相对路径
	// 在这里根据你的需求进行权限验证逻辑
	// 可以检查用户的身份、访问权限等

	// 假设所有用户都有重命名文件或目录的权限
	return true
}
