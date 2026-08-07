func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}
	sort.Ints(nums)

	var backtrack func(i int)
	backtrack = func(i int){
		if i == len(nums){
			tmp := make([]int, len(subset))
			copy(tmp, subset)
			res = append(res, tmp)
			return
		}
		subset = append(subset, nums[i])
		backtrack(i+1)
		subset = subset[:len(subset)-1]
		// Skip the duplicate number
		for i+1 < len(nums) && nums[i] == nums[i+1]{
			i++
		}
		backtrack(i+1)
	}
	backtrack(0)
	return res
}