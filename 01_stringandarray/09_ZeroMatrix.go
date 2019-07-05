package stringandarray

//ZeroMatrix 将矩阵中存在0的行和列，全部设置成0
func ZeroMatrix(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := 0; i < len(row); i++ {
		if row[i] {
			ZeroRow(matrix, i)
		}
	}
	for j := 0; j < len(col); j++ {
		if col[j] {
			ZeroCol(matrix, j)
		}
	}
}

//ZeroRow zero a row
func ZeroRow(matrix [][]int, row int) {
	for i := 0; i < len(matrix[row]); i++ {
		matrix[row][i] = 0
	}
}

//ZeroCol zero a colume
func ZeroCol(matrix [][]int, col int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][col] = 0
	}
}
