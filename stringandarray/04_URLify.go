package stringandarray

//将字符串中的空格替换成“%20”

func URLify(s string) string {
	if len(s) == 0 {
		return s
	}

	l := len(s)
	space := 0
	for i := 0; i < l; i++ {
		if s[i] == ' ' {
			space++
		}
	}
	newS := make([]byte, l+2*space)
	loc := len(newS) - 1
	i := l - 1

	for ; i >= 0; i-- {
		if s[i] == ' ' {
			newS[loc] = '0'
			newS[loc-1] = '2'
			newS[loc-2] = '%'
			loc -= 3
		} else {
			newS[loc] = s[i]
			loc--
		}
	}
	return string(newS)
}
