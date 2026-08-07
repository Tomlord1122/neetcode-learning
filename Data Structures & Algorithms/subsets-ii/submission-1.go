func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	subset := []int{}
	res := [][]int{}

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
		for i+1 < len(nums) && nums[i] == nums[i+1]{
			i++
		}
		backtrack(i+1)
	}
	backtrack(0)
	return res
}
