package paramonov_2

// SolutionUniq -- количество уникальных чисел в Массиве А
func SolutionUniq(A []int) int {
	uniq := make(map[int]bool, len(A))

	for i := 0; i < len(A); i++ {
		uniq[A[i]] = true
	}

	return len(uniq)
}
