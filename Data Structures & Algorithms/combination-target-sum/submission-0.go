func combinationSum(nums []int, target int) [][]int {
    
	// choose (stick into current index) or not choose (jump into next index)
	res := [][]int{}
	ans := []int{}
	var backtrack func(i int, value int)
	backtrack = func(i int, value int){
		if value == 0{
			tmp := make([]int, len(ans))
			copy(tmp, ans)
			res = append(res, tmp)
			return 
		}
		if i == len(nums) || value < 0{
			return
		}
		ans = append(ans, nums[i])
		backtrack(i, value-nums[i])
		// skip the current number
		ans = ans[:len(ans)-1]
		backtrack(i+1, value)
	}
	backtrack(0, target)
	return res
}
