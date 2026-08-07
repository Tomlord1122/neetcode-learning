func findDuplicate(nums []int) int {
    numCount := make(map[int]int)
	for _, num := range nums{
		if numCount[num] > 0{
			return num
		}
		numCount[num]++
	}
	return -1
}
