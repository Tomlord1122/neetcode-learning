func sortedSquares(nums []int) []int {
	var maxVal int
	for _, num := range nums{
		maxVal = max(maxVal, abs(num))
	}
	bucket := make([]int, maxVal+1)
	for _, num := range nums{
		bucket[abs(num)]++
	}

	res := []int{}

	for i := 0; i <= maxVal; i++{
		for j := 0; j < bucket[i]; j++{
			res = append(res, i*i)
		}
	}
	return res
}


func abs(x int) int{
	if x < 0{
		return -x
	}
	return x
}