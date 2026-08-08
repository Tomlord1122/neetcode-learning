func combinationSum2(candidates []int, target int) [][]int {
	cur := []int{}
	res := [][]int{}
	sort.Ints(candidates)

	var dfs func(i int, sum int)
	dfs = func(i int, sum int){
		if sum == target{
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		if i == len(candidates) || sum > target{
			return
		}
		cur = append(cur, candidates[i])
		dfs(i+1, sum+candidates[i])
		cur = cur[:len(cur)-1]
		skip := 1
		for i + skip < len(candidates) && candidates[i] == candidates[i+skip]{
			skip++
		}
		dfs(i+skip, sum)
	}
	dfs(0, 0)
	return res
}
