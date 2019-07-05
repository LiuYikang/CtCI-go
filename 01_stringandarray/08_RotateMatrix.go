package stringandarray

// RotateMatrix 矩阵顺时针旋转90度
func RotateMatrix(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	n := len(matrix)

	for layer := 0; layer < n/2; layer++ {
		last := n - 1 - layer
		for i := layer; i < last; i++ {
			offset := i - layer
			tmp := matrix[layer][i]
			matrix[layer][i] = matrix[last-offset][layer]
			matrix[last-offset][layer] = matrix[last][last-offset]
			matrix[last][last-offset] = matrix[i][last]
			matrix[i][last] = tmp
		}
	}
}
