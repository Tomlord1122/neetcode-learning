func subsets(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}

	var dfs func(i int)
	dfs = func(i int){
		// base case
		if i == len(nums){
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		cur = append(cur, nums[i])
		dfs(i+1)
		cur = cur[:len(cur)-1]
		dfs(i+1)
	}

	dfs(0)
	return res
}
