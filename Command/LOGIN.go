package Command

import (
	"fmt"
	"ftp_go/models"
	"strings"
)

func HandleUSER(dialog *models.WorkSpace, arguments []string) []byte {
	username := arguments[0]
	//todo 验证逻辑：用户名是否存在
	dialog.Usr = username
	return []byte("331 User OK\r\n")
}
func HandlePASS(dialog *models.WorkSpace, arguments []string) []byte {
	password := arguments[0]
	if !checkCredentials(dialog.Usr, password) {
		// 登录验证失败，向客户端发送错误消息并关闭连接
		return []byte("530 Login incorrect.\r\n")

	}
	dialog.Status = true
	return []byte("230 Login successful.\r\n")

}

// 验证登录
func Authenticate(dialog *models.WorkSpace) error {
	// 向客户端发送登录提示
	_, err := dialog.CommandConn.Write([]byte("220 Please enter your username:\r\n"))
	if err != nil {
		return err
	}
	// 读取客户端的用户名
	username, err := dialog.Reader.ReadString('\n')
	if err != nil {
		return err
	}
	username = strings.TrimSpace(username)

	// 向客户端发送进一步验证的提示
	_, err = dialog.CommandConn.Write([]byte("331 Please enter your password:\r\n"))
	if err != nil {
		return err
	}

	// 读取客户端的密码
	password, err := dialog.Reader.ReadString('\n')
	if err != nil {
		return err
	}
	password = strings.TrimSpace(password)
	username = strings.Replace(username, "USER ", "", -1)
	password = strings.Replace(password, "PASS ", "", -1)
	// 执行登录验证逻辑
	if !checkCredentials(username, password) {
		// 登录验证失败，向客户端发送错误消息并关闭连接
		_, err = dialog.CommandConn.Write([]byte("530 Login incorrect.\r\n"))
		if err != nil {
			return err
		}
		return fmt.Errorf("登录失败：用户名或密码错误")
	}

	// 登录验证成功，向客户端发送登录成功的消息
	_, err = dialog.CommandConn.Write([]byte("230 Login successful.\r\n"))
	if err != nil {
		return err
	}

	return nil
}

// 检查用户名和密码
func checkCredentials(username, password string) bool {
	// todo 设置校验逻辑
	// 在这里进行实际的用户名和密码验证逻辑
	// 例如，从数据库或配置文件中验证用户凭据

	// 假设用户名为 "admin"，密码为 "password"
	// 这里只是一个示例，你需要根据实际情况进行修改
	if username == "admin" && password == "password" {
		//dialog.Usr = username
		return true
	}

	return false
}
