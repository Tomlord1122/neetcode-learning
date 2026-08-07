func missingNumber(nums []int) int {
	miss := 0
	for _, num := range nums{
		miss = miss ^ num
	}
	for i := 0; i <= len(nums); i++{
		miss = miss ^ i
	}
	return miss
}
