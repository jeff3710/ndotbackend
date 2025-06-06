package util

import "strconv"

func StringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return num, nil
}

// StringToInt64 将字符串转换为 int64 类型
func StringToInt64(s string) (int64, error) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

// StringToInt32 将字符串转换为 int32 类型
func StringToInt32(s string) (int32, error) {
	num, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(num), nil
}
