package stringandarray

// IsPalindromePerm 判断一个字符串是不是可以是一个回文串的重排列
func IsPalindromePerm(s string) bool {
	counts := make(map[byte]int)
	l := len(s)

	for i := 0; i < l; i++ {
		counts[s[i]]++
	}

	odd := l%2 == 1
	seenodd := false
	for _, val := range counts {
		if val%2 == 1 {
			if !seenodd && odd {
				seenodd = true
			} else {
				return false
			}
		}
	}
	return true
}
