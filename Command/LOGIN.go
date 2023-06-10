package Command

import (
	"ftp_go/Cryption"
	"ftp_go/config"
	"ftp_go/models"
	"path/filepath"
)

func HandleUSER(dialog *models.WorkSpace, arguments []string) []byte {
	username := arguments[0]
	//判断用户是否存在
	if !config.Configs.IsSet("user." + username) {
		return []byte("530 Login incorrect.\r\n")
	}
	//验证通过
	dialog.Usr = username
	return []byte("331 User OK\r\n")
}
func HandlePASS(dialog *models.WorkSpace, arguments []string) []byte {
	password := arguments[0]
	if !checkCredentials(dialog.Usr, password) || dialog.Usr == "" {
		// 登录验证失败，向客户端发送错误消息并关闭连接
		return []byte("530 Login incorrect.\r\n")

	}
	dialog.Status = true
	dialog.Dir = filepath.Join(dialog.Dir, dialog.Usr)
	return []byte("230 Login successful.\r\n")

}

// 检查用户名和密码
func checkCredentials(username, password string) bool {
	if password == Cryption.Decode(config.Configs.GetString("user."+username+".password"), config.Configs.GetString("user."+username+".key")) {
		return true
	}

	return false
}
