package paramonov_1

import (
	"strconv"
	"strings"
)

// SolutionBinaryGap -- Поиск самого большого повторения числа 0 в 2сс
func SolutionBinaryGap(N int) int {
	result := strings.Split(strconv.FormatInt(int64(N), 2), "")
	var max = 0
	var flags = 0
	for _, value := range result {
		if value == "0" {
			flags++
		} else {
			if max < flags {
				max = flags
			}
			flags = 0
		}
	}
	return max
}
