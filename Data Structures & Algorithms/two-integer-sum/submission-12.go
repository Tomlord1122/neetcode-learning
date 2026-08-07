func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // vlaue -> index
	for i, v := range nums{
		if idx, exist :=  numMap[target-v]; exist{
			return []int{idx, i}
		}
		numMap[v] = i
	}
	return []int{}
}
