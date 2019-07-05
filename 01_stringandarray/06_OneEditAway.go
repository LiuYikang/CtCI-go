package stringandarray

// OneEditAway 判断两个字符串是否可以通过一次动作（替换/删除/增加）变成相同
func OneEditAway(str1, str2 string) bool {
	l1, l2 := len(str1), len(str2)
	dp := make([][]int, l1+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, l2+1)
	}

	for i := 0; i <= l1; i++ {
		dp[i][0] = i
	}

	for i := 0; i <= l2; i++ {
		dp[0][i] = i
	}

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if str1[i] == str2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j], min(dp[i][j+1], dp[i+1][j])) + 1
			}
		}
	}

	if dp[l1][l2] <= 1 {
		return true
	}
	return false
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
