func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	cur := []int{}
	sort.Ints(candidates)
	var backtrack func(i, sum int)
	backtrack = func(i, sum int){
		if sum == 0{
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		if i == len(candidates) || sum < 0{
			return
		}
		// choose
		cur = append(cur, candidates[i])
		backtrack(i+1, sum-candidates[i])
		// not choose
		cur = cur[:len(cur)-1]
		for i+1 < len(candidates) && candidates[i] == candidates[i+1]{
			i++
		}
		backtrack(i+1, sum)
	}
	backtrack(0, target)
	return res
}