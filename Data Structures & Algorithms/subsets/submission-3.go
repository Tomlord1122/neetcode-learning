func subsets(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}

	var helper func(i int)
	helper = func(i int){
		if i >= len(nums){
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return 
		}
		cur = append(cur, nums[i])
		helper(i+1)
		cur = cur[:len(cur)-1]
		helper(i+1)
	}
	helper(0)
	return res
}
