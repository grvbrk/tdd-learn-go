package slices

func Sum(arr []int) int {
	var total int
	for _, num := range arr {
		total += num
	}

	return total
}

func SumAll(nSlices ...[]int) []int {
	result := []int{}
	for _, arr := range nSlices {
		result = append(result, Sum(arr))
	}
	return result
}

func Return2DSlice(nSlices ...[]int) [][]int {
	return nSlices
}
