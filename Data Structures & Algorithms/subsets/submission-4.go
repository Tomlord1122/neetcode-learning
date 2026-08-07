func subsets(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}

	var backtrack func(i int) 
	backtrack = func(i int){
		if i == len(nums){
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		// choose the cur num
		cur = append(cur, nums[i])
		backtrack(i+1)
		// skip the cur num
		cur = cur[:len(cur)-1]
		backtrack(i+1)
	}

	backtrack(0)
	return res
}



