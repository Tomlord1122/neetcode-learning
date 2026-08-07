func canPartition(nums []int) bool {
    target := 0
	for _, v := range nums{
		target += v
	}
	if target % 2 != 0{
		return false
	}
	target = target / 2

	var dfs func(i int, target int) bool
	dfs = func(i int, target int) bool{
		if target == 0{
			return true
		}
		if i >= len(nums) || target < 0{
			return false
		}
		return dfs(i+1, target) || dfs(i+1, target-nums[i])
	}
	return dfs(0, target)
}


