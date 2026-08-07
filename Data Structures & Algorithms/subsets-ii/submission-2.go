func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}
	sort.Ints(nums)
	var backtrack func(i int)
	backtrack = func(i int){
		if i == len(nums){
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return 
		}
		// choose
		cur = append(cur, nums[i])
		backtrack(i+1)
		// not choose and skip duplicate
		cur = cur[:len(cur)-1]
		for i+1 < len(nums) && nums[i] == nums[i+1]{
			i++
		}
		backtrack(i+1)
	}
	backtrack(0)
	return res
}
