func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int) // value, index pair
	for idx2, num := range nums{
		if idx1, exist := numMap[target-num]; exist{
			return []int{idx1, idx2}
		}
		numMap[num] = idx2
	}
	return []int{}
}
