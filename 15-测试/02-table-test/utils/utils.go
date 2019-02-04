package utils

func Sum(arr ...int) int {
	var result int
	for _, value := range arr {
		result += value
	}
	return result
}
