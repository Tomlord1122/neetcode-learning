func singleNumber(nums []int) int {
	target := 0
	for _, num := range nums{
		target = target ^ num
	}
	return target
}
