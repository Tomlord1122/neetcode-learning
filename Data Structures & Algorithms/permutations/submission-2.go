func permute(nums []int) [][]int {
	if len(nums) == 0{
		return [][]int{{}}
	}

	perms := permute(nums[1:])
	res := [][]int{}
	for _, p := range perms{
		for i := 0; i <= len(p); i++{
			// left
			left := append([]int{}, p[:i]...)
			// right 
			right := append([]int{nums[0]}, p[i:]...)
			// merge
			merge := append(left, right...)
			// res
			res = append(res, merge)
		}
	}
	return res
}
