package server

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "ftp_go",
		Usage: "一个简单的CLI应用程序",
		Commands: []*cli.Command{
			//{
			//	Name:  "login",
			//	Usage: "登录到FTP服务器",
			//	Action: func(c *cli.Context) error {
			//		host := c.String("host")
			//		port := c.Int("port")
			//		username := c.String("username")
			//		password := c.String("password")
			//
			//		// 创建 FTP 客户端
			//		ftpClient, err := ftp.Connect(host + ":" + strconv.Itoa(port))
			//		if err != nil {
			//			return err
			//		}
			//
			//		// 登录到 FTP 服务器
			//		err = ftpClient.Login(username, password)
			//		if err != nil {
			//			return err
			//		}
			//		defer ftpClient.Logout()
			//
			//		// 执行 FTP 操作，这里仅为示例
			//		// 可以根据需求执行其他的 FTP 操作，例如上传、下载文件等
			//
			//		fmt.Println("成功登录到 FTP 服务器")
			//
			//		return nil
			//	},
			//	Flags: []cli.Flag{
			//		&cli.StringFlag{
			//			Name:     "host",
			//			Usage:    "FTP 服务器主机名",
			//			Required: true,
			//		},
			//		&cli.IntFlag{
			//			Name:  "port",
			//			Usage: "FTP 服务器端口号",
			//			Value: 21,
			//		},
			//		&cli.StringFlag{
			//			Name:     "username",
			//			Aliases:  []string{"u"},
			//			Usage:    "FTP 服务器用户名",
			//			Required: true,
			//		},
			//		&cli.StringFlag{
			//			Name:     "password",
			//			Aliases:  []string{"p"},
			//			Usage:    "FTP 服务器密码",
			//			Required: true,
			//			Hidden:   true,
			//		},
			//	},
			//},
			{
				Name:  "start",
				Usage: "启动 Ftp server",
				Action: func(c *cli.Context) error {
					Start()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
