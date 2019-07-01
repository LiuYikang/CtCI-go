package stringandarray

//IsUnique 判断一个字符串里的字符是不是全部唯一，假设所有字符都是ASCII
func IsUnique(str string) bool {
	if len(str) > 256 {
		return false
	}

	exit := make([]bool, 256)

	for _, c := range str {
		if exit[int(c)] {
			return false
		}
		exit[int(c)] = true
	}
	return true
}
