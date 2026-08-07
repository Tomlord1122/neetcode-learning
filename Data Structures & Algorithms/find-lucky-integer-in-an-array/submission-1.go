func findLucky(arr []int) int {
	numCount := make(map[int]int)
	for _, num := range arr{
		numCount[num]++
	}
	res := -1
	for val, freq := range numCount{
		if val == freq{
			res = max(res, val)
		}
	}
	return res
}
