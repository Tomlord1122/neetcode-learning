func combinationSum(nums []int, target int) [][]int {
    cur := []int{}
	res := [][]int{}

	var backtrack func(i, sum int)
	backtrack = func(i, sum int){
		if i == len(nums) || sum < 0{
			return
		}
		if sum == 0{
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		cur = append(cur, nums[i])
		backtrack(i, sum-nums[i])
		cur = cur[:len(cur)-1]
		backtrack(i+1, sum)
	}
	backtrack(0, target)
	return res
}
