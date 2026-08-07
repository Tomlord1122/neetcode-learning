func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int) // val, idx
	for idx2, val := range nums{
		if idx1, exist := numMap[target-val]; exist{
			return []int{idx1, idx2}
		}
		numMap[val] = idx2
	}
	return []int{}
}
