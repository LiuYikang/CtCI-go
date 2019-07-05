package stringandarray

//Reverse 反转字符串
func Reverse(str string) string {
	if len(str) == 0 {
		return str
	}

	chs := []byte(str)
	low, high := 0, len(chs)-1
	for low < high {
		chs[low], chs[high] = chs[high], chs[low]
		low++
		high--
	}
	return string(chs)
}
