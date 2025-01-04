package utils

func CalcSliceSum(slice []int) int {
	var total int
	for _, value := range slice {
		total += value
	}
	return total
}
