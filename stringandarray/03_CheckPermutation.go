package stringandarray

//判断一个字符串重新排列之后能不能是另一个字符串

func CheckPermutation(s, t string) bool {
	l1, l2 := len(s), len(t)

	if l1 == 0 && l2 == 0 {
		return true
	}

	if l1 != l2 {
		return false
	}

	m := make(map[byte]int)
	for i := 0; i < l1; i++ {
		m[s[i]]++
	}

	for i := 0; i < l2; i++ {
		m[t[i]]--
		if m[t[i]] < 0 {
			return false
		}
	}
	return true
}
