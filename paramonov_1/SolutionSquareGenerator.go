package paramonov_1

// SolutionSquareGenerator -- Генератор квадратов
func SolutionSquareGenerator(start int, n int) []int {
	var result = make([]int, n)

	for i := 0; i < n; i++ {
		var endCircle int = start + i
		result[i] = endCircle * endCircle
	}

	return result
}
