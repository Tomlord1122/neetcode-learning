func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	cur := []int{}

	var dfs func(i, sum int)
	dfs = func(i, sum int){
		if sum == 0{
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		if i >= len(candidates) || sum < 0{
			return
		}
		cur = append(cur, candidates[i])
		dfs(i+1, sum-candidates[i])
		cur = cur[:len(cur)-1]
		for i + 1 < len(candidates) && candidates[i] == candidates[i+1]{
			i++
		}
		dfs(i+1, sum)
	}
	dfs(0, target)
	return res
}
