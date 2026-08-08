func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	cur := []int{}
	res := [][]int{}

	var bt func(i int)
	bt = func(i int){
		if i == len(nums){
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		cur = append(cur, nums[i])
		bt(i+1)
		cur = cur[:len(cur)-1]
		skip := 1
		for i+skip < len(nums) && nums[i] == nums[i+skip]{
			skip++
		}
		bt(i+skip)
	}
	bt(0)
	return res
}
