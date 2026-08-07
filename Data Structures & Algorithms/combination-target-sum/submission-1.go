func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}
	cur := []int{}
	// when we skip the current num, we should jump into next index
	// otherwise, we should remain the index
	var backtrack func(idx int, sum int)
	backtrack = func(idx int, sum int){
		if idx == len(nums) || sum < 0{
			return
		}
		if sum == 0{
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		cur = append(cur, nums[idx])
		backtrack(idx, sum-nums[idx])
		cur = cur[:len(cur)-1]
		backtrack(idx+1, sum)
	}
	// call the backtrack function
	backtrack(0, target)
	return res
}
