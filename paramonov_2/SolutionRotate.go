package paramonov_2

// SolutionRotate -- Сдвиг массива A, по K
func SolutionRotate(A []int, k int) []int {
	var rotateArray []int = make([]int, len(A))

	pos := len(A) - k

	for i := 0; i < k; i++ {
		rotateArray[i] = A[pos+i]
	}

	for i := 0; i < pos; i++ {
		rotateArray[k+i] = A[i]
	}

	return rotateArray
}
