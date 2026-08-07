func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}
	sort.Ints(nums)

	var bt func(i int, cur []int)
	bt = func(i int, cur []int){
		if i >= len(nums){
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		cur = append(cur, nums[i])
		bt(i+1, cur)
		cur = cur[:len(cur)-1]
		for i + 1 < len(nums) && nums[i] == nums[i+1]{
			i++
		}
		bt(i+1, cur)
	}
	bt(0, cur)
	return res
}
