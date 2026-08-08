func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	cur := []int{}
	res := [][]int{}
	var backtrack func(i int)
	backtrack = func(i int){
		if i == len(nums){
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return

		}
		cur = append(cur, nums[i])
		backtrack(i+1)
		count := 1
		for i+count < len(nums) && nums[i+count] == nums[i]{
			count++
		}
		cur = cur[:len(cur)-1]
		backtrack(i+count)
	}
	backtrack(0)
	return res
}
// if we sort nums first
// then we can use a backtracking function to solve this
// when we skip the current val, we should also skip the following duplicated num