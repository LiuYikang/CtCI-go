package stringandarray

import (
	"strconv"
)

// CompressStr 压缩字符串，string aabcccccaaa would become a2blc5a3
func CompressStr(str string) string {
	var compStr string
	count := 1
	i := 1
	for ; i < len(str); i++ {
		if str[i-1] == str[i] {
			count++
		} else {
			compStr = compStr + string(str[i-1]) + strconv.Itoa(count)
			count = 1
		}
	}
	compStr = compStr + string(str[i-1]) + strconv.Itoa(count)

	if len(compStr) > len(str) {
		return str
	}
	return compStr
}
