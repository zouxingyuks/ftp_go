package Command

import (
	"fmt"
	"ftp_go/Response"
	"ftp_go/models"
	"os/exec"
	"strings"
	"time"
)

func HandleListCommand(dialog *models.WorkSpace, arguments []string) []byte {
	dir := dialog.Dir
	if len(arguments) > 0 {
		dir = arguments[0]
	}
	//dir = config.Configs.GetString("dir.root") + "/" + dialog.Usr + dir
	dir = dialog.BasicDir + dir
	//todo 测试
	fmt.Println("list dir: ", dir)
	cmd := exec.Command("bash", "-c", fmt.Sprintf("ls -l %s | tail -n +2", dir))
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		Response.Send(dialog.CommandConn, []byte("500 List Error\r\n"), dialog.TransferType)
		return nil
	}
	Response.Send(dialog.CommandConn, []byte("200 List OK\r\n"), dialog.TransferType)
	Response.Send(dialog.DataConn, append(output, []byte("\r\n")...), dialog.TransferType)
	dialog.DataConn.Close()
	return nil
}

// 格式化文件列表
func formatFileList(fileList string) string {
	lines := strings.Split(fileList, "\n")
	var formattedList strings.Builder

	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 9 {
			permissions := fields[0]
			linkCount := fields[1]
			owner := fields[2]
			group := fields[3]
			size := fields[4]
			month := fields[5]
			day := fields[6]
			if len(day) < 2 {
				day = " " + day
			}
			timeStr := fields[7]
			filename := fields[8]

			// 解析日期字符串
			dateStr := month + " " + day + " " + timeStr
			date, err := time.Parse("1月 2 15:04", dateStr)
			if err != nil {
				// 解析失败，跳过该行
				continue
			}
			// 格式化日期为所需的格式
			formattedDate := date.Format("01 02 15:04")
			//formattedLine := fmt.Sprintf("%s %s %s %s %s %s %s\n", permissions, linkCount, size, owner, group, formattedDate, filename)
			formattedLine := fmt.Sprintf("%s %s %s %s %s %s %s\n", permissions, linkCount, owner, group, size, formattedDate, filename)
			formattedList.WriteString(formattedLine)
		}
	}

	return formattedList.String()
}
