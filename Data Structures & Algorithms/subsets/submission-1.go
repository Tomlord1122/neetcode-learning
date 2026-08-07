func subsets(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	cur := []int{}
	var bt func(i int)
	bt = func(i int){
		if i == n{
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		cur = append(cur, nums[i])
		bt(i+1)
		cur = cur[:len(cur)-1]
		bt(i+1)
	}
	bt(0)
	return res
}