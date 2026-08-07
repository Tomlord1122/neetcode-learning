func singleNumber(nums []int) int {
	miss := 0
	for _, num := range nums{
		miss = miss ^ num
	}
	return miss
}
