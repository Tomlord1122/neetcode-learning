func missingNumber(nums []int) int {
	n := len(nums)
	miss := 0
	for _, num := range nums{
		miss = miss ^ num
	}
	for i := 0; i <= n; i++{
		miss = miss ^ i
	}
	return miss
}
