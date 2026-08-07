func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int) // (val, idx)
	for i, num := range nums{
		if idx1,exist :=numMap[target-num]; exist{
			return []int{idx1, i}
		}
		numMap[num] = i
	}
	return []int{}
}
