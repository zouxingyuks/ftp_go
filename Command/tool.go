package Command

import "strconv"

// 辅助函数：将字符串转换为整数
func toInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}
