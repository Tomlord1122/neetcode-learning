func majorityElement(nums []int) int {
    majorFreq := len(nums) / 2
	numCount := make(map[int]int)
	for _, val := range nums{
		numCount[val]++
		if numCount[val] > majorFreq{
			return val
		}
	}
	return -1
}
