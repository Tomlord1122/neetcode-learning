func subsetXORSum(nums []int) int {
	cur := []int{}
	res := 0
	var backtrack func(i int)
	backtrack = func(i int){
		if i == len(nums){
			v := 0
			for _, val := range cur{
				v ^= val
			}
			res += v
			return
		}
		cur = append(cur, nums[i])
		backtrack(i+1)
		cur = cur[:len(cur)-1]
		backtrack(i+1)
	}
	backtrack(0)
	return res
}
